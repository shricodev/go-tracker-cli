package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)
// GetUserInput retrieves user input either from the provided arguments or by
// reading from a reader stream, such as stdin. If arguments are passed, they are
// joined into a single string. If no arguments are provided, the function reads
// from the reader until a newline character is encountered. Returns an error if
// the input is empty after trimming whitespace.
func GetUserInput(reader io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	// When the user sends the data from a pipe command
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("There was an error reading the input %w", err)
	}

	text := strings.TrimSpace(scanner.Text())

	if len(text) == 0 {
		return "", errors.New("Cannot add an empty Tracker")
	}

	return text, nil
}
