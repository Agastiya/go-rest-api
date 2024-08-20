package Validator

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

func Validate(mystruct interface{}) error {
	validate = validator.New()
	err := validate.Struct(mystruct)
	if err != nil {
		var errorValidate = "error when validation, error:"
		if _, ok := err.(*validator.InvalidValidationError); ok {
			errorValidate += fmt.Sprintf("%s ,", err.Error())
		}
		errValidator := err.(validator.ValidationErrors)
		for index, err := range errValidator {
			if index+1 == len(errValidator) {
				errorValidate += err.Field() + " " + err.Tag()
			} else {
				errorValidate += err.Field() + " " + err.Tag() + ","
			}
		}
		return errors.New(errorValidate)
	}
	return nil
}

func ValidateEnum(enum []string, data interface{}) bool {
	for _, dataEnum := range enum {
		if dataEnum == data {
			return true
		}
	}
	return false
}
