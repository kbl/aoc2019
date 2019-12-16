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

func TestDigitValueAfterIteration(t *testing.T) {
	type TC struct {
		digits       []int
		numberLength int
		index        int
		iteration    int
		expected     int
	}

	tcs := []TC{
		{[]int{1, 2, 3, 4, 5, 6}, 6, 3, 1, 5},
		{[]int{1, 2, 3, 4, 5, 6}, 6, 2, 1, 8},
		{[]int{1, 2, 3, 4, 5, 6}, 6, 2, 2, 0},
		{[]int{1, 2, 3, 4, 5, 6}, 12, 7, 2, 0},
		{[]int{1, 2, 3, 4, 5, 6}, 12, 6, 2, 1},
		{[]int{1, 2, 3, 4, 5, 6}, 12, 3, 2, 6},
		{[]int{1, 2, 3, 4, 5, 6}, 12, 0, 2, 8},
	}

	for _, tc := range tcs {
		got := digitValueAfterIteration(tc.digits, tc.numberLength, tc.iteration, tc.index)
		if got != tc.expected {
			t.Errorf("digitValueAfterIteration(%v, %d, %d, %d) = %d, want %d", tc.digits, tc.numberLength, tc.iteration, tc.index, got, tc.expected)
		}
	}
}

func TestAdvancedPhase(t *testing.T) {
	tcs := [][]string{
		{"03036732577212944063491565474664", "84462026"},
		{"02935109699940807407585447034323", "78725270"},
		{"03081770884921959731165446850517", "53553731"},
	}

	for _, tc := range tcs {
		got := AdvancedPhase(100, tc[0])
		if got != tc[1] {
			t.Errorf("Phase(100, %s) = %s, want %s", tc[0], got, tc[1])
		}
	}
}
