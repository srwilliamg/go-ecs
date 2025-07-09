package customError

type CustomError struct {
	Message string   `json:"message"`
	Detail  []string `json:"detail"`
}

func NewCustomError(message string, detail []string) *CustomError {
	return &CustomError{Message: message}
}

func (e *CustomError) Error() string {
	return e.Message
}
