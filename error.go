package main

import "fmt"

type CustomError struct {
	Code    int
	Message string
}

func NewCustomError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func ExampleFunctionThatReturnsError() error {
	return NewCustomError(404, "Resource not found")
}
