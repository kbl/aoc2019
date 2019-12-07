package main

import (
	"testing"
)

type testCase struct {
	memory       []int
	phaseOptions []int
	expected     int
}

var testCases = []testCase{
	{[]int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}, []int{0, 1, 2, 3, 4}, 43210},
	{[]int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}, []int{0, 1, 2, 3, 4}, 54321},
	{[]int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}, []int{0, 1, 2, 3, 4}, 65210},
}

func TestSomething(t *testing.T) {
	for _, tc := range testCases {
		is := NewIntcodeSequence(tc.memory, tc.phaseOptions)
		got := is.Run()

		if got != tc.expected {
			t.Errorf("is.Run() = %d, want %d", got, tc.expected)
		}
	}
}
