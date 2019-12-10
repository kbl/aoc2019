package main

import (
	"reflect"
	"testing"
)

type testCase struct {
	s        string
	expected map[Cord]int
}

var testCases = []testCase{
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
###.##.####.##.#..##`, map[Cord]int{{11, 13}: 222}},
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

func TestNewSpace(t *testing.T) {
	for _, tc := range testCases {
		s := NewSpace(tc.s)
		for c, expected := range tc.expected {
			if s.Visible(c) != expected {
				t.Errorf("s.Visible(%v) = %d, wants %d!", c, s.Visible(c), expected)
			}
		}
	}
}
