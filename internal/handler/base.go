package handler

import (
	"Wordle/internal/response"

	"github.com/go-playground/validator/v10"
)

func parseValidationErrors(err error) []response.ValidationError {
	var errors []response.ValidationError
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			var ve response.ValidationError
			ve.Loc = []string{e.Field()}
			ve.Msg = e.Error()
			ve.Type = e.Tag()
			errors = append(errors, ve)
		}
	}
	return errors
}
