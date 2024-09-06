package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockDatabaseObject(t *testing.T) {
	// create an instance
	testObj := new(MockDatabaseObject)
	testObj.Data = make(map[string]string)

	// set value to a key
	testObj.Set("name", "Shane")

	// get value from a key
	value, err := testObj.Get("name")

	if err != nil {
		assert.Fail(t, err.Error())
	} else {
		assert.Equal(t, "Shane", value)
	}
}
