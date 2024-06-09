package handler

import (
	"fmt"
	"net/http"
	"strconv"

	cnst "github.com/LostArrows27/snippetbox/pkg/const"
	htmlParse "github.com/LostArrows27/snippetbox/pkg/html-parse"
	ipaddress "github.com/LostArrows27/snippetbox/pkg/ip-address"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	ipaddress.LogRequestIP("/", r)

	htmlParse.ExecuteHTML(w, cnst.HomeBase, cnst.HomeHTMLLists)
}

func CreateSnippetHanlder(w http.ResponseWriter, r *http.Request) {
	ipaddress.LogRequestIP("/snippet/create", r)
	w.Write([]byte("Create snippet"))
}

func ViewSnippetHandler(w http.ResponseWriter, r *http.Request) {
	ipaddress.LogRequestIP("/snippet/view", r)

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 0 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func StaticFileHanlder(w http.ResponseWriter, r *http.Request) {
	fileServer := http.FileServer(http.Dir(cnst.StaticFileDir))

	handlerFile := http.StripPrefix("/static", fileServer)

	handlerFile.ServeHTTP(w, r)
}
