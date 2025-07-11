package request

import (
	"encoding/json"
)

func MarshalResponse[T any](obj T, err error) ([]byte, error) {
	var sentError error = nil
	if err != nil {
		sentError = err
	}
	res := BaseResponse(obj, sentError, nil)
	resJsonBytes, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	return resJsonBytes, err
}
