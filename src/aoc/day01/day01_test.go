package main

import (
	"testing"
)

type testCase struct {
	input, expectedFuel, expectedTotalFuel int
}

var testCases = []testCase{
	{12, 2, 2},
	{14, 2, 2},
	{1969, 654, 966},
	{100756, 33583, 50346},
}

func TestFuel(t *testing.T) {
	for _, tc := range testCases {
		got := Fuel(tc.input)

		if got != tc.expectedFuel {
			t.Errorf("Fuel(%d) = %d, want %d", tc.input, got, tc.expectedFuel)
		}
	}
}

func TestTotalFuel(t *testing.T) {
	for _, tc := range testCases {
		got := TotalFuel(tc.input)

		if got != tc.expectedTotalFuel {
			t.Errorf("TotalFuel(%d) = %d, want %d", tc.input, got, tc.expectedTotalFuel)
		}
	}
}
