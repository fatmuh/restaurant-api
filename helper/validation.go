package helper

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

func FormatValidationError(err error) map[string]string {
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return map[string]string{"error": "Invalid request"}
	}

	errors := make(map[string]string)
	for _, fieldError := range validationErrors {
		field := fieldError.Field()
		tag := fieldError.Tag()

		var message string
		switch tag {
		case "required":
			message = field + " is required"
		case "min":
			message = field + " must be at least " + fieldError.Param() + " characters"
		case "max":
			message = field + " must be at most " + fieldError.Param() + " characters"
		case "email":
			message = field + " must be a valid email address"
		// Add more custom error messages as needed
		default:
			message = field + " is not valid"
		}

		// Convert field name to lowercase with underscores (e.g., "Name" to "name")
		field = strings.ToLower(field)
		errors[field] = message
	}

	return errors
}
