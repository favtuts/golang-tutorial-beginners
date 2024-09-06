// create mock objects by implementing interfaces directly
package main

type MockDatabaseImplInterface struct {
	GetFunc func(key string) (string, error)
	SetFunc func(key, value string) error
}

func (db *MockDatabaseImplInterface) Get(key string) (string, error) {
	return db.GetFunc(key)
}

func (db *MockDatabaseImplInterface) Set(key, value string) error {
	return db.SetFunc(key, value)
}
