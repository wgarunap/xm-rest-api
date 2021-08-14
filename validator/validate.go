package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/wgarunap/xm-rest-api/domain"
)

type Validator struct {
	validate *validator.Validate
}

func (v Validator) Validate(_struct interface{}) error {
	err := v.validate.Struct(_struct)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return err
		}

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}
		return err
	}
	return nil
}

func NewValidator() domain.Validator {
	return &Validator{validate: validator.New()}
}
