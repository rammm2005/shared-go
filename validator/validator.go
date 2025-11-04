package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ValidateStruct validates a struct using tags like `validate:"required"`
func ValidateStruct(s interface{}) error {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	var messages []string
	for _, e := range err.(validator.ValidationErrors) {
		messages = append(messages, fmt.Sprintf("%s is %s", e.Field(), e.Tag()))
	}

	return fmt.Errorf(strings.Join(messages, ", "))
}
