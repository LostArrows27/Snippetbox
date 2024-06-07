package main

import (
	"log"
	"net/http"

	"github.com/LostArrows27/snippetbox/pkg/env"
	ipaddress "github.com/LostArrows27/snippetbox/pkg/ip-address"
	"github.com/LostArrows27/snippetbox/pkg/rest"
)

func main() {
	// 0. load ENV
	env.LoadEnv(".env")
	port := env.GetEnv("PORT")

	// 1. Server IPv4 address
	ips, err := ipaddress.GetServerIP()
	if err != nil {
		log.Println("Error:", err)
		return
	}
	log.Println("Server IPs:", ips[0])

	// 2. configre route
	restMux := rest.RestAPI{
		MUX: http.NewServeMux(),
	}
	restMux.Get("/", home, "fixed")
	restMux.Get("/snippet/view", viewSnippetHandler)
	restMux.Post("/snippet/create", createSnippetHanlder)

	// 3. run server
	log.Printf("Starting server on %v", port)
	err = http.ListenAndServe(":"+port, restMux.MUX)

	if err != nil {
		log.Fatal(err)
	}
}
