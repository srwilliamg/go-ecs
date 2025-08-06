package validator

import (
	"fmt"
	"io"
	"srwilliamg/app/v1/internal/application/request"
	log "srwilliamg/app/v1/internal/interfaces/logger"

	"github.com/go-playground/validator/v10"
)

func DecodeAndValidateBody[T any](logger *log.Logger, validate *validator.Validate, body io.ReadCloser, dto *T) []string {
	if err := request.DecodeBody(body, dto); err != nil {
		errorMessage := err.Error()
		(*logger).Error(errorMessage)
		return []string{errorMessage}
	}

	validationErrorsSlice := ValidateStruct(validate, *dto)

	if validationErrorsSlice != nil {
		(*logger).Error("Invalid body")
		return validationErrorsSlice
	}

	return nil
}

func ValidateStruct[T any](validate *validator.Validate, data T) []string {
	err := validate.Struct(data)
	var errorSlice []string
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return errorSlice
		}

		for _, err := range err.(validator.ValidationErrors) {
			reqBodyError := fmt.Sprintf("Field %s failed validation: %s", err.Field(), err.Tag())
			fmt.Println(reqBodyError)
			errorSlice = append(errorSlice, reqBodyError)
		}
	}

	return errorSlice
}
