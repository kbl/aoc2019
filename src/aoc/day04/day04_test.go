package main

import (
	"testing"
)

type testCase struct {
	password       int
	valid1, valid2 bool
}

var testCases = []testCase{
	{111111, true, false},
	{111122, true, true},
	{112233, true, true},
	{123444, true, false},
	{223450, false, false},
	{123789, false, false},
}

func TestValidPassword(t *testing.T) {
	for _, tc := range testCases {
		got := ValidPassword(tc.password)
		if got != tc.valid1 {
			t.Errorf("ValidPassword(%d) = %t, want %t", tc.password, got, tc.valid1)
		}
	}
}

func TestValidPassword2(t *testing.T) {
	for _, tc := range testCases {
		got := ValidPassword2(tc.password)
		if got != tc.valid2 {
			t.Errorf("ValidPassword2(%d) = %t, want %t", tc.password, got, tc.valid2)
		}
	}
}
