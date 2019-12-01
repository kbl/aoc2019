package day01

import (
	"testing"
)

var fuelTestCases = [][]int{
	[]int{12, 2},
	[]int{14, 2},
	[]int{1969, 654},
	[]int{100756, 33583},
}

var totalFuelTestCases = [][]int{
	[]int{12, 2},
	[]int{1969, 966},
	[]int{100756, 50346},
}

func TestFuel(t *testing.T) {
	for _, tc := range fuelTestCases {
		input := tc[0]
		expected := tc[1]

		got := Fuel(input)

		if got != expected {
			t.Errorf("Fuel(%d) = %d, want %d", input, got, expected)
		}
	}
}

func TestTotalFuel(t *testing.T) {
	for _, tc := range totalFuelTestCases {
		input := tc[0]
		expected := tc[1]

		got := TotalFuel(input)

		if got != expected {
			t.Errorf("TotalFuel(%d) = %d, want %d", input, got, expected)
		}
	}
}
