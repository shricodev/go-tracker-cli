package utils

import (
	"fmt"
	"os"
)

type InvalidIndexError struct {
	Index int
}

// Error implements the error interface by providing a human-readable description
// of the invalid index error. The returned string contains the invalid index value.
func (e *InvalidIndexError) Error() string {
	return fmt.Sprintf("Invalid index: %d", e.Index)
}

// HandleGenericError prints the error message to stderr and exits with code 1 if err is not nil.
func HandleGenericError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
