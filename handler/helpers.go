package handler

import (
	"errors"
	"net/http"

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
