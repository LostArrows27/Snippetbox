package handler

import (
	"net/http"
	"time"

	"github.com/LostArrows27/snippetbox/internal/models"
)

type templateData struct {
	CurrentYear     int
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	Title           string
	Form            any
	Flash           string
	IsAuthenticated bool
	UserData        models.UserData
}

func (app *Application) newTemplateData(r *http.Request) *templateData {
	// remove the flash from the request -> not show next time

	return &templateData{
		CurrentYear:     time.Now().Year(),
		Flash:           app.SessionManager.PopString(r.Context(), "flash"),
		IsAuthenticated: app.IsAuthenticated(r),
		UserData:        app.GetAuthenticatedUserData(r),
	}
}
