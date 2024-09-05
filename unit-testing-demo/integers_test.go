// integers_test.go
package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	arg1 int
	arg2 int
	want int
}

func TestMultiply(t *testing.T) {
	got := Multiply(2, 3)
	want := 6

	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5

	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestMultiplyTableDriven(t *testing.T) {
	cases := []testCase{
		{2, 3, 6},
		{10, 5, 50},
		{-8, -3, 24},
		{0, 9, 0},
		{-7, 6, -42},
	}

	for _, tc := range cases {
		got := Multiply(tc.arg1, tc.arg2)
		if tc.want != got {
			t.Errorf("Expected '%d', but got '%d'", tc.want, got)
		}
	}
}

func TestMultiplySubtests(t *testing.T) {
	cases := []testCase{
		{2, 3, 6},
		{10, 5, 50},
		{-8, -3, 24},
		{0, 9, 0},
		{-7, 6, -42},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d*%d=%d", tc.arg1, tc.arg2, tc.want), func(t *testing.T) {
			got := Multiply(tc.arg1, tc.arg2)
			if tc.want != got {
				t.Errorf("Expected '%d', but got '%d'", tc.want, got)
			}
		})
	}
}

func TestAddTableDriven(t *testing.T) {
	cases := []testCase{
		{1, 1, 2},
		{7, 5, 12},
		{-19, -3, -22},
		{-1, 8, 7},
		{-12, 0, -12},
	}

	for _, tc := range cases {
		got := Add(tc.arg1, tc.arg2)
		if tc.want != got {
			t.Errorf("Expected '%d', but got '%d'", tc.want, got)
		}
	}
}
