package main

import (
	"testing"
)

type testCase struct {
	input    int
	expected int
}

var testCases = []testCase{}

func TestSomething(t *testing.T) {
	for _, tc := range testCases {
		got := tc.expected

		if got != tc.expected {
			t.Errorf("dummy(%d) = %d, want %d", tc.input, got, tc.expected)
		}
	}
}
