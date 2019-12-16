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

func TestPhaseMultiplier(t *testing.T) {
	tcs := [][]int{
		{1, 0, 1},
		{1, 1, 0},
		{1, 2, -1},
		{1, 3, 0},
		{1, 4, 1},
		{1, 5, 0},
		{1, 6, -1},

		// 0 0 1 1 0 0 -1 -1
		{2, 0, 0},
		{2, 1, 1},
		{2, 2, 1},
		{2, 3, 0},
		{2, 4, 0},
		{2, 5, -1},
		{2, 6, -1},
		{2, 7, 0},
		{2, 8, 0},
		{2, 9, 1},

		// 0 0 0 1 1 1 0 0 0 -1 -1 -1
		{3, 0, 0},
		{3, 1, 0},
		{3, 2, 1},
		{3, 3, 1},
		{3, 4, 1},
		{3, 5, 0},
		{3, 6, 0},
		{3, 7, 0},
		{3, 8, -1},
		{3, 9, -1},
		{3, 10, -1},
		{3, 11, 0},
		{3, 12, 0},
		{3, 13, 0},
		{3, 14, 1},
		{3, 15, 1},
	}
	for _, tc := range tcs {
		phase := tc[0]
		index := tc[1]
		expected := tc[2]

		got := phaseMultiplier(phase, index)

		if got != expected {
			t.Errorf("phaseMultiplier(%d, %d) = %d, want %d", phase, index, got, expected)
		}
	}
}

func TestPhase(t *testing.T) {
	type TC struct {
		phases       int
		in, expected string
	}

	tcs := []TC{
		{1, "12345678", "48226158"},
		{2, "12345678", "34040438"},
		{3, "12345678", "03415518"},
		{4, "12345678", "01029498"},
		{100, "80871224585914546619083218645595", "24176176"},
	}

	for _, tc := range tcs {
		got := Phase(tc.phases, tc.in)
		if got != tc.expected {
			t.Errorf("Phase(%d, %s) = %s, want %s", tc.phases, tc.in, got, tc.expected)
		}
	}
}
