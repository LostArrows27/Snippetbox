package main

import (
	"fmt"
	"net/http"
	"strconv"

	htmlParse "github.com/LostArrows27/snippetbox/pkg/html-parse"
	ipaddress "github.com/LostArrows27/snippetbox/pkg/ip-address"
)

func home(w http.ResponseWriter, r *http.Request) {
	ipaddress.LogRequestIP("/", r)

	htmlParse.ExecuteHTML(w, "./ui/html/page/home.html")

}

func createSnippetHanlder(w http.ResponseWriter, r *http.Request) {
	ipaddress.LogRequestIP("/snippet/create", r)
	w.Write([]byte("Create snippet"))
}

func viewSnippetHandler(w http.ResponseWriter, r *http.Request) {
	ipaddress.LogRequestIP("/snippet/view", r)

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 0 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}
