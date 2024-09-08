package main

import (
	"github.com/pkg/errors"
)

func X() error {
	return errors.Errorf("Could not write to file (pkg.errors)")
}

func CustomErrorPkgErrors() error {
	return X()
}
