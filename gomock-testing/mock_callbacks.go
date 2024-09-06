package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func ProcessData(data string, f func(string) string) string {
	result := f(data)
	// do some processing with result
	return result
}

func TestProcessData(t *testing.T) {
	mockFunc := func(data string) string {
		return "mocked result"
	}
	result := ProcessData("input data", mockFunc)
	assert.Equal(t, "mocked result", result)
}
