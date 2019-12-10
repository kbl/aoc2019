package main

import (
	"reflect"
	"testing"
)

type visibilityTestCase struct {
	s        string
	expected map[Cord]int
}

var visibilityTestCases = []visibilityTestCase{
	{`.#..#
.....
#####
....#
...##`, map[Cord]int{{3, 4}: 8, {4, 4}: 7},
	},
	{`.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`, map[Cord]int{{11, 13}: 210}},
	{`###..#########.#####.
.####.#####..####.#.#
.###.#.#.#####.##..##
##.####.#.###########
###...#.####.#.#.####
#.##..###.########...
#.#######.##.#######.
.#..#.#..###...####.#
#######.##.##.###..##
#.#......#....#.#.#..
######.###.#.#.##...#
####.#...#.#######.#.
.######.#####.#######
##.##.##.#####.##.#.#
###.#######..##.#....
###.##.##..##.#####.#
##.########.#.#.#####
.##....##..###.#...#.
#..#.####.######..###
..#.####.############
..##...###..#########`, map[Cord]int{{11, 11}: 221}},
}

func TestVisible(t *testing.T) {
	for _, tc := range visibilityTestCases {
		s := NewSpace(tc.s)
		for c, expected := range tc.expected {
			if s.Visible(c) != expected {
				t.Errorf("s.Visible(%v) = %d, wants %d!", c, s.Visible(c), expected)
			}
		}
	}
}

func TestPickVisible(t *testing.T) {
	cords := Cords{Cord{0, 0}, Cord{1, 1}, Cord{2, 2}}
	tcs := map[Cord]Cords{
		Cord{1, 1}: Cords{Cord{0, 0}, Cord{2, 2}},
		Cord{0, 0}: Cords{Cord{1, 1}},
		Cord{2, 2}: Cords{Cord{1, 1}},
		Cord{7, 1}: Cords{},
	}

	for c, expected := range tcs {
		if !reflect.DeepEqual(pickVisible(c, cords), expected) {
			t.Errorf("pickVisible(%v) = %v, wants %v!", c, pickVisible(c, cords), expected)
		}
	}
}

func TestNewFunc(t *testing.T) {
	cords := map[Func]Cords{
		Func{2, 0}:    Cords{{0, 0}, {1, 2}},
		Func{1, 0}:    Cords{{0, 0}, {1, 1}},
		Func{0.5, 0}:  Cords{{0, 0}, {2, 1}},
		Func{0, 0}:    Cords{{0, 0}, {1, 0}},
		Func{-0.5, 0}: Cords{{0, 0}, {-2, 1}},
		Func{-1, 0}:   Cords{{0, 0}, {-1, 1}},
		Func{-2, 0}:   Cords{{0, 0}, {-1, 2}},
	}
	for expected, c := range cords {
		got := NewFunc(c[0], c[1])
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("NewFunc(%v, %v) = %v, wants %v!", c[0], c[1], got, expected)
		}
	}
}

func TestVaporize(t *testing.T) {
	sp := `.#....#####...#..
##...##.#####..##
##...#...#.#####.
..#.....#...###..
..#.#.....#....##`
	s := NewSpace(sp)

	// for c, expected := range tc.expected {
	// 	if s.Visible(c) != expected {
	// 		t.Errorf("s.Visible(%v) = %d, wants %d!", c, s.Visible(c), expected)
	// 	}
	// }
	s.Vaporize(Cord{8, 3}, 36)
}
