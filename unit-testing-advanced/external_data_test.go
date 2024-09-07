package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetBase64Encoding(t *testing.T) {
	cases := []string{"iphone", "android"}

	for _, v := range cases {
		t.Run(v, func(t *testing.T) {
			b, err := os.ReadFile(filepath.Join("testdata", v+".png"))
			if err != nil {
				t.Fatal(err)
			}

			expected, err := os.ReadFile(filepath.Join("testdata", v+"_base64.txt"))
			if err != nil {
				t.Fatal(err)
			}

			got := getBase64Encoding(b)

			if string(expected) != got {
				t.Fatalf("Expected output to be: '%s', but got: '%s'", string(expected), got)
			}
		})
	}
}
