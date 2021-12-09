package errors

import (
	"fmt"
	"net/http"
)

//ErrInvalidState struct
type ErrInvalidState struct {
	State string
}

//NewErrInvalidState constructor.
func NewErrInvalidState(state string) *ErrInvalidState {
	err := &ErrInvalidState{State: state}
	return err
}

//Error taking the cause.
func (e *ErrInvalidState) Error() string {
	if e == nil {
		return "<nil>"
	}

	return fmt.Sprintf("Invalid state %v. [Status code %v]", e.State, http.StatusBadRequest)
}
