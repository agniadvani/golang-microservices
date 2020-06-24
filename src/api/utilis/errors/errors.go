package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ApiError interface {
	Status() int
	Message() string
	Error() string
}

type apiErr struct {
	Astatus  int    `json:"status"`
	Amessage string `json:"message"`
	Aerror   string `json:"error,omitempty"`
}

func (s *apiErr) Status() int {
	return s.Astatus
}

func (s *apiErr) Message() string {
	return s.Amessage
}

func (s *apiErr) Error() string {
	return s.Aerror
}
func NewApiErrFromBytes(body []byte) (ApiError, error) {
	var result apiErr
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("invalid json body")
	}
	return &result, nil
}
func NewApiErr(statusCode int, message string) ApiError {
	return &apiErr{
		Astatus:  statusCode,
		Amessage: message,
	}
}
func NewInternalServerError(message string) ApiError {
	return &apiErr{
		Astatus:  http.StatusInternalServerError,
		Amessage: message,
	}
}
func NewBadRequestError(message string) ApiError {
	return &apiErr{
		Astatus:  http.StatusBadRequest,
		Amessage: message,
	}
}

func NewNotFoundError(message string) ApiError {
	return &apiErr{
		Astatus:  http.StatusNotFound,
		Amessage: message,
	}
}
