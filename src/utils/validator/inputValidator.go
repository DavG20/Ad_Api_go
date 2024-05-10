package validator

import (
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

func ValidateTime(fl validator.FieldLevel) bool {
	dateStr := fl.Field().Interface().(time.Time).Format("2006-01-02")

	_, err := time.Parse("2006-01-02", dateStr)
	return err == nil
}

var Validate *validator.Validate

// Register the custom validation function
func init() {
	Validate = validator.New()
	Validate.RegisterValidation("ValidateTime", ValidateTime)
}

// ValidationError represents a validation error.
type ValidationError struct {
	Field   string
	Message string
}

// ValidateInputs validates the given dataset using the go-playground/validator package.
// It returns a slice of validation errors.
func ValidateInputs(dataset interface{}) (bool, []ValidationError) {
	// Validate = validator.New()

	err := Validate.Struct(dataset)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			// This error indicates a bug in the validator itself, panic in this case
			return false, []ValidationError{
				{
					Field:   "validator",
					Message: "validator error"},
			}

		}

		var errors []ValidationError
		reflected := reflect.ValueOf(dataset)

		for _, err := range err.(validator.ValidationErrors) {
			field, _ := reflected.Type().FieldByName(err.StructField())
			name := field.Tag.Get("json")
			if name == "" {
				name = strings.ToLower(err.StructField())
			}

			if err.Tag() == "ValidateTime" {
				errors = append(errors, ValidationError{
					Field:   name,
					Message: getErrorMessage(name, err),
				})
			} else {
				errors = append(errors, ValidationError{
					Field:   name,
					Message: getErrorMessage(name, err),
				})
			}
		}

		return false, errors
	}

	return true, nil
}

// getErrorMessage returns a human-readable error message for the validation error.
func getErrorMessage(name string, err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "The " + name + " is required"

	case "email":
		return "The " + strings.Split(name, ",")[0] + " should be a valid email"

	case "gte":
		return "The " + strings.Split(name, ",")[0] + " must be greater than 0"

	case "startswith":
		return "The " + strings.Split(name, ",")[0] + " should be start with " + err.Param()
	case "min":
		return "The length of " + name + " should be equal to or greater than " + err.Param()
	case "ltefield":
		return "The " + name + " should be less than or equal to " + err.Param()
	case "iso4217":
		return "The " + strings.Split(name, ",")[0] + " should be a valid ISO 4217 currency code"
	case "eqfield":
		return "The " + name + " should be equal to " + err.Param()
	case "gtefield":
		return "The " + name + " should be greater than or equal to " + err.Param()
	case "ValidateTime":
		return "The " + strings.Split(name, ",")[0] + " has to be a valid date"
	case "iso3166_1_alpha3":
		return "The " + strings.Split(name, ",")[0] + " should be a valid ISO 3166-1 alpha-3 country code"
	case "url":
		return "The " + strings.Split(name, ",")[0] + " should be a valid URL"
	case "boolean":
		return "The " + strings.Split(name, ",")[0] + " should be a valid boolean"
	default:
		return "The " + name + " is invalid"
	}
}
