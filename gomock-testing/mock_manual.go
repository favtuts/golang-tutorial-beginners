// Manually creating mock objects involves creating a struct
// that implements the same interface as the real object,
// but with mock data and behavior.
package main

import "errors"

type MockDatabaseObject struct {
	Data map[string]string
}

func (db *MockDatabaseObject) Get(key string) (string, error) {
	value, ok := db.Data[key]
	if !ok {
		return "", errors.New("key not found")
	}
	return value, nil
}

func (db *MockDatabaseObject) Set(key, value string) error {
	db.Data[key] = value
	return nil
}
