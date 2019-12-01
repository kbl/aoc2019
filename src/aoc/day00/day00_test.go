package day00

import (
	"testing"
)

var testCases = [][]int{
	[]int{2, 2},
}

func TestFuel(t *testing.T) {
	for _, tc := range testCases {
		input := tc[0]
		expected := tc[1]

		got := expected

		if got != expected {
			t.Errorf("dummy(%d) = %d, want %d", input, got, expected)
		}
	}
}
