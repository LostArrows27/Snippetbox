package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from SnippetBox"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	port := 4000
	log.Printf("Starting server on %v", port)

	err := http.ListenAndServe(":4000", mux)

	if err != nil {
		log.Fatal(err)
	}
}
