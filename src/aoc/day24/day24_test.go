package main

import (
	"fmt"
	"strings"
	"testing"
)

type testCase struct {
	input    int
	expected int
}

var testCases = []testCase{}

func TestSomething(t *testing.T) {
	input := `....#
#..#.
#..##
..#..
#....`
	grid := parse(strings.Split(input, "\n"))
	fmt.Println(str(grid))
	fmt.Println(first(grid))
}
