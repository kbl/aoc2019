package main

import (
	"testing"
)

type testCase struct {
	input        string
	steps        int
	energy       int
	periodLength int
}

var testCases = []testCase{
	testCase{`<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`,
		10, 179, 2772},
	testCase{`<x=-8, y=-10, z=0>
<x=5, y=5, z=10>
<x=2, y=-7, z=3>
<x=9, y=-8, z=-3>`,
		100, 1940, 4686774924},
	testCase{`<x=14, y=4, z=5>
<x=12, y=10, z=8>
<x=1, y=7, z=-10>
<x=16, y=-5, z=3>`,
		1000, 6423, 327636285682704},
}

func TestEnergyAfter(t *testing.T) {
	for _, tc := range testCases {
		s := NewSpace(tc.input)
		got := s.EnergyAfter(tc.steps)

		if got != tc.energy {
			t.Errorf("s.EnergyAfter(%d) = %d, want %d", tc.steps, got, tc.energy)
		}
	}
}

func TestPeriodLength(t *testing.T) {
	for _, tc := range testCases {
		s := NewSpace(tc.input)
		got := s.PeriodLength()

		if got != tc.periodLength {
			t.Errorf("s.PeriodLength() = %d, want %d", got, tc.periodLength)
		}
	}
}
