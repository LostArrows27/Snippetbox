package main

import (
	"log"
	"net/http"

	"github.com/LostArrows27/snippetbox/env"
	ipaddress "github.com/LostArrows27/snippetbox/helper/ip-address"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ip := ipaddress.GetIP(r)

	log.Printf("Request to / : from %s", ip)
	w.Write([]byte("Hello from SnippetBox"))
}

func createSnippetHanlder(w http.ResponseWriter, r *http.Request) {
	ip := ipaddress.GetIP(r)
	log.Printf("Request to /snippet/create : from %s", ip)
	w.Write([]byte("Create snippet"))
}

func viewSnippetHandler(w http.ResponseWriter, r *http.Request) {
	ip := ipaddress.GetIP(r)
	log.Printf("Request to /snippet/view : from %s", ip)
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
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", viewSnippetHandler)
	mux.HandleFunc("/snippet/create", createSnippetHanlder)

	// 3. run server
	log.Printf("Starting server on %v", port)
	err = http.ListenAndServe(":"+port, mux)

	if err != nil {
		log.Fatal(err)
	}
}
