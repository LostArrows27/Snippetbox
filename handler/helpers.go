package handler

import (
	"errors"
	"net/http"

	"github.com/LostArrows27/snippetbox/internal/models"
	"github.com/go-playground/form/v4" // New import
)

func (app *Application) decodePostForm(r *http.Request, dst any) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}
	err = app.FormDecoder.Decode(dst, r.PostForm)
	if err != nil {
		var invalidDecoderError *form.InvalidDecoderError
		if errors.As(err, &invalidDecoderError) {
			panic(err)
		}
		return err
	}
	return nil
}

func (app *Application) IsAuthenticated(r *http.Request) bool {
	return app.SessionManager.Exists(r.Context(), "authenticatedUserID")
}

func (app *Application) GetAuthenticatedUserData(r *http.Request) models.UserData {
	if app.SessionManager.Exists(r.Context(), "authenticatedUserID") {
		return models.UserData{
			ID:   app.SessionManager.GetInt(r.Context(), "authenticatedUserID"),
			Name: app.SessionManager.GetString(r.Context(), "authenticatedUserName"),
		}
	}

	return models.UserData{}
}
