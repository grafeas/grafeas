package errors

import "fmt"

// AppError conveys statused errors
type AppError struct {
	Err        string
	StatusCode int
}

// Error is the consolidated error message output for an AppError and satisfies
// the builtin error interface.
func (e AppError) Error() string {
	return fmt.Sprintf("[%d] %s", e.StatusCode, e.Err)
}
