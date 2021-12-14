package errors

import (
	"fmt"
)

//ErrInvalidState struct
type ErrInvalidState struct {
	State string
}

//NewErrInvalidState constructor.
func NewErrInvalidState(state string) *ErrInvalidState {
	return &ErrInvalidState{State: state}
}

//Error taking the cause.
func (e *ErrInvalidState) Error() string {
	if e == nil {
		return "<nil>"
	}

	return fmt.Sprintf("Invalid state %v", e.State)
}
