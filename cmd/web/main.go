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

	// 1. erver IPv4 address
	ips, err := ipaddress.GetServerIP()
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("Server IPs: %v", ips[0])
	logger.Info("Starting server on port: %v", port)

	// 2. configure application global variables + dependency
	app := &handler.Application{
		ErrorLog: *logger.ErrorLogger(),
		InfoLog:  *logger.InfoLogger(),
	}

	// 3 configure rest API to pass in app router
	restMux := rest.RestAPI{
		MUX: http.NewServeMux(),
	}

	// 4. configure server + run server
	srv := &http.Server{
		Addr:     ":" + port,
		ErrorLog: logger.ErrorLogger().Logger,
		Handler:  app.RoutesHandler(restMux),
	}
	err = srv.ListenAndServe()

	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
