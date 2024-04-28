package user

import (
	"errors"
	"fmt"
)

var ErrFirstNameRequired = errors.New("first name is required")
var ErrLastNameRequired = errors.New("second name is required")
var ErrEmailRequired = errors.New("email is required")

type ErrNotFound struct {
	ID uint64
}

func (err ErrNotFound) Error() string {
	return fmt.Sprintf("user with id %d doesn't exist", err.ID)
}
