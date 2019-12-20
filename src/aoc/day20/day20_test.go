package main

import (
	"reflect"
	"testing"
)

type testCase struct {
	maze            string
	vertices, edges int
	entrances       map[Vertex][]cord
	shortestPath    int
}

//        11111111112222222222333333
//2345678901234567890123456789012345

var testCases = []testCase{
	testCase{
		maze: `         A           
         A           
  #######.#########  
  #######.........#  
  #######.#######.#  
  #######.#######.#  
  #######.#######.#  
  #####  B    ###.#  
BC...##  C    ###.#  
  ##.##       ###.#  
  ##...DE  F  ###.#  
  #####    G  ###.#  
  #########.#####.#  
DE..#######...###.#  
  #.#########.###.#  
FG..#########.....#  
  ###########.#####  
             Z       
             Z       `,
		vertices: 5,
		edges:    8,
		entrances: map[Vertex][]cord{
			"AA": []cord{{9, 2}},
			"BC": []cord{{2, 8}, {9, 6}},
			"DE": []cord{{2, 13}, {6, 10}},
			"FG": []cord{{2, 15}, {11, 12}},
			"ZZ": []cord{{13, 16}},
		},
		shortestPath: 23,
	},
	testCase{
		maze: `                   A               
                   A               
  #################.#############  
  #.#...#...................#.#.#  
  #.#.#.###.###.###.#########.#.#  
  #.#.#.......#...#.....#.#.#...#  
  #.#########.###.#####.#.#.###.#  
  #.............#.#.....#.......#  
  ###.###########.###.#####.#.#.#  
  #.....#        A   C    #.#.#.#  
  #######        S   P    #####.#  
  #.#...#                 #......VT
  #.#.#.#                 #.#####  
  #...#.#               YN....#.#  
  #.###.#                 #####.#  
DI....#.#                 #.....#  
  #####.#                 #.###.#  
ZZ......#               QG....#..AS
  ###.###                 #######  
JO..#.#.#                 #.....#  
  #.#.#.#                 ###.#.#  
  #...#..DI             BU....#..LF
  #####.#                 #.#####  
YN......#               VT..#....QG
  #.###.#                 #.###.#  
  #.#...#                 #.....#  
  ###.###    J L     J    #.#.###  
  #.....#    O F     P    #.#...#  
  #.###.#####.#.#####.#####.###.#  
  #...#.#.#...#.....#.....#.#...#  
  #.#####.###.###.#.#.#########.#  
  #...#.#.....#...#.#.#.#.....#.#  
  #.###.#####.###.###.#.#.#######  
  #.#.........#...#.............#  
  #########.###.###.#############  
           B   J   C               
           U   P   P               `,
		vertices: 12,
		edges:    18,
		entrances: map[Vertex][]cord{
			"AA": []cord{{19, 2}},
			"AS": []cord{{32, 17}, {17, 8}},
			"BU": []cord{{11, 34}, {26, 21}},
			"CP": []cord{{19, 34}, {21, 8}},
			"DI": []cord{{2, 15}, {8, 21}},
			"JO": []cord{{2, 19}, {13, 28}},
			"JP": []cord{{15, 34}, {21, 28}},
			"LF": []cord{{32, 21}, {15, 28}},
			"QG": []cord{{32, 23}, {26, 17}},
			"VT": []cord{{32, 11}, {26, 23}},
			"YN": []cord{{2, 23}, {26, 13}},
			"ZZ": []cord{{2, 17}},
		},
		shortestPath: 58,
	},
}

func TestSomething(t *testing.T) {
	for _, tc := range testCases {
		g := NewGraph(tc.maze)
		if len(g.Edges) != tc.edges {
			t.Errorf("len(g.Edges) = %d, want %d", len(g.Edges), tc.edges)
		}
		if len(g.Vertices) != tc.vertices {
			t.Errorf("len(g.Vertices) = %d, want %d", len(g.Vertices), tc.vertices)
		}
		if !reflect.DeepEqual(g.entrances, tc.entrances) {
			t.Errorf("g.entrances = %v, want %v", g.entrances, tc.entrances)
		}
		if g.ShortestPath("AA", "ZZ") != tc.shortestPath {
			t.Errorf("g.ShortestPath(AA, ZZ) = %d, want %d", g.ShortestPath("AA", "ZZ"), tc.shortestPath)
		}
	}
}
