package request

import (
	"reflect"
	customError "srwilliamg/app/v1/internal/application/custom-error"
	log "srwilliamg/app/v1/internal/interfaces/logger"
)

type controllerResponse[T any] struct {
	Data  T     `json:"data"`
	Error error `json:"error"`
}

func BaseResponse[T any](Data T, err error, logger *log.Logger) *controllerResponse[T] {
	if err != nil && logger != nil {
		(*logger).Error("error", log.Err(err))
	}

	if err != nil && reflect.TypeOf(err).String() == "*errors.errorString" {
		return &controllerResponse[T]{Error: &customError.CustomError{Message: err.Error()}}
	}

	return &controllerResponse[T]{Data, err}
}
