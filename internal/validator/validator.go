package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func New() *Validator {
	v := validator.New(validator.WithRequiredStructEnabled())

	// use json tag in error messages instead of struct field
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.Split(fld.Tag.Get("json"), ",")[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return &Validator{
		validate: v,
	}
}

func (v *Validator) Validate(i any) error {
	return v.validate.Struct(i)
}

func FormatErrors(err error) map[string]string {
	errors := map[string]string{}

	if err == nil {
		return nil
	}

	for _, err := range err.(validator.ValidationErrors) {
		field := err.Field()

		switch err.Tag() {
		case "required":
			errors[field] = "is required"
		case "email":
			errors[field] = "must be a valid email"
		case "min":
			errors[field] = "is too short"
		case "max":
			errors[field] = "is too long"
		default:
			errors[field] = "invalid value"
		}
	}

	return errors
}
