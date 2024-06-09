package main

import (
	"net/http"

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

	// 2. configre route
	restMux := rest.RestAPI{
		MUX: http.NewServeMux(),
	}
	restMux.Get("/static/", handler.StaticFileHanlder)
	restMux.Get("/", handler.HomeHandler, "fixed")
	restMux.Get("/snippet/view", handler.ViewSnippetHandler)
	restMux.Post("/snippet/create", handler.CreateSnippetHanlder)

	// 3. run server
	logger.Info("Starting server on %v", port)
	err = http.ListenAndServe(":"+port, restMux.MUX)

	if err != nil {
		logger.Error(err)
	}
}
