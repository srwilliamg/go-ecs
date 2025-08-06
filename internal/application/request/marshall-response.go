package request

import (
	"encoding/json"
)

func MarshalResponse[T any](obj T) ([]byte, error) {
	res := BaseResponse(obj, nil)

	resJsonBytes, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	return resJsonBytes, err
}
