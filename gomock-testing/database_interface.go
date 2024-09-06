package main

type Database interface {
	Get(key string) (string, error)
	Set(key, value string) error
}
