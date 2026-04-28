package validator

import "github.com/go-playground/validator/v10"

type Handler struct {
}

type Validator struct {
	v *validator.Validate
}

func New() *Validator {
	return &Validator{v: validator.New()}
}
