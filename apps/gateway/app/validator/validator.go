package validator

import "github.com/go-playground/validator/v10"

type MyValidator struct {
	*validator.Validate
}

func GetValidator() *MyValidator {
	return &MyValidator{
		validator.New(),
	}
}
