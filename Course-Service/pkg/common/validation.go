package common

import (
	"github.com/asaskevich/govalidator"
	"net/http"
)

func CatchErrorValidation(err error) error {
	if errs, ok := err.(govalidator.Errors); ok {
		// Collect all validation error messages
		var validationErrors []string
		for _, e := range errs {
			validationErrors = append(validationErrors, e.Error())
		}
		return NewCustomError(err, http.StatusUnsupportedMediaType, validationErrors[0])
	}
	return err
}
