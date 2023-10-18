package e

import "fmt"

// Wrap "оборачивает ошибку"
func Wrap(msg string, err error) error {
	return fmt.Errorf("%s: %w", msg, err)
}
