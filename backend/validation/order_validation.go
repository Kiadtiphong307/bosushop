package validation

import (
	"github.com/go-playground/validator/v10"
)	

type OrderInput struct {
	ProductID   uint   `json:"product_id" validate:"required"`
	CouponCode  string `json:"coupon_code"` // optional
}

func ValidateOrderInput(input OrderInput) map[string]string {
	err := Validate.Struct(input) // ✅ เรียกจาก validator.go
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		errors[e.Field()] = e.ActualTag()
	}
	return errors
}
