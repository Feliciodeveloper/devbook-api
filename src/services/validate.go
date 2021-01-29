package services

import "github.com/go-playground/validator/v10"

var validate *validator.Validate


func ValidateStruct(param interface{})error{
	validate = validator.New()
	return validate.Struct(param)
}
func ValidateEmail(param string)error {
	validate = validator.New()
	if err := validate.Var(param,"email"); err != nil {
		return err
	}
	return nil
}
