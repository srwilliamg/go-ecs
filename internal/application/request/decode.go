package request

import (
	"encoding/json"
	"io"
)

func DecodeBody[T any](body io.ReadCloser, dto *T) error {
	defer body.Close()
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&dto)

	return err
}
