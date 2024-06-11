package main

import (
	"net/http"
	"os"

	"github.com/LostArrows27/snippetbox/handler"
	"github.com/LostArrows27/snippetbox/pkg/env"
	ipaddress "github.com/LostArrows27/snippetbox/pkg/ip-address"
	"github.com/LostArrows27/snippetbox/pkg/logger"
	"github.com/LostArrows27/snippetbox/pkg/rest"
)

func main() {
	// 0. load ENV
	env.LoadEnv(".env")
	port := env.GetEnv("PORT")

	// 1. Server IPv4 address
	ips, err := ipaddress.GetServerIP()
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("Server IPs: %v", ips[0])

	// 2. configure application global variables + dependency
	app := &handler.Application{
		ErrorLog: *logger.ErrorLogger(),
		InfoLog:  *logger.InfoLogger(),
	}

	// 2. configre route
	restMux := rest.RestAPI{
		MUX: http.NewServeMux(),
	}
	restMux.Get("/static/", app.HomeHandler)
	restMux.Get("/", app.HomeHandler, "fixed")
	restMux.Get("/snippet/view", app.ViewSnippetHandler)
	restMux.Post("/snippet/create", app.CreateSnippetHanlder)

	// 3. configure server + run server
	srv := &http.Server{
		Addr:     ":" + port,
		ErrorLog: logger.ErrorLogger().Logger,
		Handler:  restMux.MUX,
	}
	err = srv.ListenAndServe()

	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
