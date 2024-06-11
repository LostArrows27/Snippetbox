package htmlParse

import (
	"net/http"
	"text/template"
)

func ParseHTML(w http.ResponseWriter, path []string) (*template.Template, error) {
	ts, err := template.ParseFiles(path...)

	if err != nil {
		return nil, err
	}

	return ts, nil

}

func ExecuteHTML(w http.ResponseWriter, ts *template.Template, base string, path []string) error {

	err := ts.ExecuteTemplate(w, base, nil)
	if err != nil {
		return err
	}

	return nil
}
