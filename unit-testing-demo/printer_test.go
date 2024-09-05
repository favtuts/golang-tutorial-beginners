package main

import (
	"bytes"
	"strings"
	"testing"
)

// printer_test.go
func TestPrint(t *testing.T) {
	var buf bytes.Buffer

	text := "Hello, World!"

	Print(text, &buf)

	got := strings.TrimSpace(buf.String())

	if got != text {
		t.Errorf("Expected output to be: %s, but got: %s", text, got)
	}
}
