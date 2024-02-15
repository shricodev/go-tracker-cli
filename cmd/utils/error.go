package utils

import (
	"fmt"
	"os"
)

type InvalidIndexError struct {
	Index int
}

func (e *InvalidIndexError) Error() string {
	return fmt.Sprintf("Invalid index: %d", e.Index)
}

func HandleGenericError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
