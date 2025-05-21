package validation

import (
	"github.com/go-playground/validator/v10"
)

type ProductInput struct {
	Name        string  `json:"name" validate:"required,min=3"`
	Description string  `json:"description"`
	ImageURL    string  `json:"image_url"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	CategoryID  uint    `json:"category_id" validate:"required"`
	Available   bool    `json:"available"`
	Stock       int     `json:"stock" validate:"required,gte=0"`
}

func ValidateProductInput(input ProductInput) map[string]string {
	validate := validator.New()
	err := validate.Struct(input)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		errors[e.Field()] = e.ActualTag()
	}
	return errors
}
