package main

import (
	"bytes"
	"testing"

	"github.com/sebdah/goldie/v2"
)

func TestGenerateTable(t *testing.T) {
	var buf bytes.Buffer

	err := generateTable(inventory, &buf)
	if err != nil {
		t.Fatal(err)
	}

	actual := buf.Bytes()

	g := goldie.New(t)
	g.Assert(t, "books", actual)
}
