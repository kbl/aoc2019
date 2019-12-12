package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	input    string
	steps    int
	expected int
}

var testCases = []testCase{
	testCase{`<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`,
		10,
		179},
}

func TestSomething(t *testing.T) {
	for _, tc := range testCases {
		s := NewSpace(tc.input)
		s.Simulate(tc.steps)
		got := s.Energy()

		if got != tc.expected {
			t.Errorf("s.Energy() = %d, want %d", got, tc.expected)
		}
	}
}
