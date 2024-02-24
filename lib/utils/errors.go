package utils

import "fmt"

func ErrNotFound(msg string) error {
	return fmt.Errorf("NOT FOUND: %s", msg)
}
