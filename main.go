package main

import (
	"log"
	"net/http"

	"github.com/LostArrows27/snippetbox/env"
	ipaddress "github.com/LostArrows27/snippetbox/helper/ip-address"
	"github.com/LostArrows27/snippetbox/rest"
)

func logRequestIP(path string, r *http.Request) {
	ip := ipaddress.GetIP(r)

	log.Printf("Request to %s : from %s", path, ip)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	logRequestIP("/", r)
	w.Write([]byte("Hello from SnippetBox"))
}

func createSnippetHanlder(w http.ResponseWriter, r *http.Request) {
	logRequestIP("/snippet/create", r)
	w.Write([]byte("Create snippet"))
}

func viewSnippetHandler(w http.ResponseWriter, r *http.Request) {
	logRequestIP("/snippet/view", r)

	w.Write([]byte("View snippet"))
}

func main() {
	// 1. load ENV
	env.LoadEnv(".env")
	port := env.GetEnv("PORT")

	// 1.5. Server IPv4 address
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

	restMux.Get("/", home)
	restMux.Get("/snippet/view", viewSnippetHandler)
	restMux.Post("/snippet/create", createSnippetHanlder)

	// 3. run server
	log.Printf("Starting server on %v", port)
	err = http.ListenAndServe(":"+port, restMux.MUX)

	if err != nil {
		log.Fatal(err)
	}
}
