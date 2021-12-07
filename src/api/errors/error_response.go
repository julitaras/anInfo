package errors

import (
	"fmt"
)

//ErrResponse struct
type ErrResponse struct {
	Err     error
	Message string
}

//NewErrResponse constructor.
func NewErrResponse(err error) *ErrResponse {
	return &ErrResponse{Err: err, Message: err.Error()}
}

//Error taking the cause.
func (e *ErrResponse) Error() string {
	if e == nil {
		return "<nil>"
	}

	return fmt.Sprintf("Cause error %v", e.Error())
}
