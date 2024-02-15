package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

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
