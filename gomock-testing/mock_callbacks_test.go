package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessData(t *testing.T) {
	mockFunc := func(data string) string {
		return "mocked result"
	}
	result := ProcessData("input data", mockFunc)
	assert.Equal(t, "mocked result", result)
}
