package appRequest

import (
	"reflect"
	customError "srwilliamg/app/v1/internal/custom-error"

	zap "go.uber.org/zap"
)

type controllerResponse[T any] struct {
	Data  T     `json:"data"`
	Error error `json:"error"`
}

func BaseResponse[T any](Data T, err error, logger *zap.Logger) *controllerResponse[T] {
	if err != nil && logger != nil {
		logger.Error("error", zap.Error(err))
	}

	if err != nil && reflect.TypeOf(err).String() == "*errors.errorString" {
		return &controllerResponse[T]{Error: &customError.CustomError{Message: err.Error()}}
	}

	return &controllerResponse[T]{Data, err}
}
