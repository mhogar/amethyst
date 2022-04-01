package common

import (
	"errors"
	"fmt"
)

// ChainError will combine the error message and the message together in an easy to read manner.
func ChainError(message string, err error) error {
	return errors.New(message + "\n\t" + err.Error())
}

func NewError(format string, args ...any) error {
	return errors.New(
		fmt.Sprintf(format, args...),
	)
}
