package rest

import "ockernuts/goroborally/models"

// Error captures the rest error code and a rendering of the error text according to the rest error model.
type Error interface {
	error

	GetResultCode() int
	GetRestError() *models.Error
}

// NewRestError creates a new error carrier, indicating a rest code
func NewRestError(code int, parentError error) Error {
	return restError{resultCode: code, parentError: parentError}
}

type restError struct {
	resultCode  int
	parentError error
}

func (r restError) GetResultCode() int {
	return r.resultCode
}

func (r restError) GetRestError() *models.Error {
	result := models.Error{Message: r.parentError.Error()}
	return &result
}

func (r restError) Error() string {
	return r.parentError.Error()
}
