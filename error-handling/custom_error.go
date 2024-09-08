package main

import "errors"

// Define an error struct
type CustomError struct {
	msg string
}

// Create a function Error() string and associate it to the struct
// This is to implement the built-in error interface
func (error *CustomError) Error() string {
	return error.msg
}

// Then create an error object using MyError struct
func CustomErrorInstance() error {
	return &CustomError{
		"File type not supported (CustomError)",
	}
}

// The newly created custom error can then be restructured to use the built-in error struct:
func NewCustomError() error {
	return errors.New("File type not supported (errors.New)")
}
