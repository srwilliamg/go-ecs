package request

import (
	log "srwilliamg/app/v1/internal/interfaces/logger"
)

type controllerResponse[T any] struct {
	Data  T     `json:"data"`
	Error error `json:"error"`
}

func BaseResponse[T any](Data T, logger *log.Logger) *controllerResponse[T] {
	if v, ok := any(Data).(error); ok && v != nil {
		return &controllerResponse[T]{Error: v}
	}

	return &controllerResponse[T]{Data, nil}
}
