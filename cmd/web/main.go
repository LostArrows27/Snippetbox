package main

import (
	"net/http"
	"os"
	"time"

	"github.com/LostArrows27/snippetbox/handler"
	"github.com/LostArrows27/snippetbox/internal/models"
	"github.com/LostArrows27/snippetbox/pkg/database"
	"github.com/LostArrows27/snippetbox/pkg/env"
	ipaddress "github.com/LostArrows27/snippetbox/pkg/ip-address"
	"github.com/LostArrows27/snippetbox/pkg/logger"
	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
)

func main() {
	// 0. load ENV + init dependency
	env.LoadEnv(".env")
	port := env.GetEnv("PORT")
	dbURL := env.GetEnv("DB_URL")
	errorLog := logger.ErrorLogger()
	infoLog := logger.InfoLogger()
	form := form.NewDecoder()
	template, err := handler.NewTemplateCache() // load all html page into template in runtime

	if err != nil {
		logger.Error(err)
	}

	// 1. connect to database
	db, err := database.OpenDB(dbURL)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// 2. configure session manager as dependency to application
	sessionManager := scs.New()
	sessionManager.Store = postgresstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour
	sessionManager.Cookie.Secure = true

	// 3. log server IPv4 address + port
	ips, err := ipaddress.GetServerIP()
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("Server IPs: %v", ips[0])
	logger.Info("Starting server on port: %v", port)

	// 4. configure application global variables + dependency
	app := &handler.Application{
		ErrorLog:       *errorLog,
		InfoLog:        *infoLog,
		Snippets:       &models.SnippetModel{DB: db},
		TemplateCache:  template,
		FormDecoder:    form,
		SessionManager: sessionManager,
	}

	// 5. configure server + run HTTPS server
	srv := &http.Server{
		Addr:     ":" + port,
		ErrorLog: errorLog.Logger,
		Handler:  app.RoutesHandler(),
	}
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")

	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
