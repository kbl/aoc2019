package main

import (
	"strings"
	"testing"
)

func TestExercise1(t *testing.T) {
	input := `....#
#..#.
#..##
..#..
#....`
	grid := parse(strings.Split(input, "\n"))
	got := exercise1(grid)
	expected := 2129920

	if got != expected {
		t.Errorf("exercise1(grid) = %d, want %d", got, expected)
	}
}

func TestExercise2(t *testing.T) {
	input := `....#
#..#.
#..##
..#..
#....`
	grid := parse(strings.Split(input, "\n"))
	got := exercise2(10, grid)
	expected := 99

	if got != expected {
		t.Errorf("exercise2(10, grid) = %d, want %d", got, expected)
	}
}
