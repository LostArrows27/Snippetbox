package main

import (
	"net/http"
	"os"

	"github.com/LostArrows27/snippetbox/handler"
	"github.com/LostArrows27/snippetbox/pkg/database"
	"github.com/LostArrows27/snippetbox/pkg/env"
	ipaddress "github.com/LostArrows27/snippetbox/pkg/ip-address"
	"github.com/LostArrows27/snippetbox/pkg/logger"
	"github.com/LostArrows27/snippetbox/pkg/rest"
)

func main() {
	// 0. load ENV + init dependency
	env.LoadEnv(".env")
	port := env.GetEnv("PORT")
	dbURL := env.GetEnv("DB_URL")
	errorLog := logger.ErrorLogger()
	infoLog := logger.InfoLogger()

	// 1. connect to database
	db, err := database.OpenDB(dbURL)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// 2. log server IPv4 address + port
	ips, err := ipaddress.GetServerIP()
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("Server IPs: %v", ips[0])
	logger.Info("Starting server on port: %v", port)

	// 3. configure application global variables + dependency
	app := &handler.Application{
		ErrorLog: *errorLog,
		InfoLog:  *infoLog,
	}

	// 4. configure rest API to pass in app router
	restMux := rest.RestAPI{
		MUX: http.NewServeMux(),
	}

	// 5. configure server + run server
	srv := &http.Server{
		Addr:     ":" + port,
		ErrorLog: errorLog.Logger,
		Handler:  app.RoutesHandler(restMux),
	}
	err = srv.ListenAndServe()

	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
