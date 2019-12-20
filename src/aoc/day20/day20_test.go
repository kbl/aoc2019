package main

import (
	"reflect"
	"testing"
)

type testCase struct {
	maze            string
	vertices, edges int
	entrances       map[Vertex]cord
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
		entrances: map[Vertex]cord{
			Vertex{"AA", outer}: cord{9, 2},
			Vertex{"BC", outer}: cord{2, 8},
			Vertex{"BC", inner}: cord{9, 6},
			Vertex{"DE", outer}: cord{2, 13},
			Vertex{"DE", inner}: cord{6, 10},
			Vertex{"FG", outer}: cord{2, 15},
			Vertex{"FG", inner}: cord{11, 12},
			Vertex{"ZZ", outer}: cord{13, 16},
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
		entrances: map[Vertex]cord{
			Vertex{"AA", outer}: cord{19, 2},
			Vertex{"AS", outer}: cord{32, 17},
			Vertex{"AS", inner}: cord{17, 8},
			Vertex{"BU", outer}: cord{11, 34},
			Vertex{"BU", inner}: cord{26, 21},
			Vertex{"CP", outer}: cord{19, 34},
			Vertex{"CP", inner}: cord{21, 8},
			Vertex{"DI", outer}: cord{2, 15},
			Vertex{"DI", inner}: cord{8, 21},
			Vertex{"JO", outer}: cord{2, 19},
			Vertex{"JO", inner}: cord{13, 28},
			Vertex{"JP", outer}: cord{15, 34},
			Vertex{"JP", inner}: cord{21, 28},
			Vertex{"LF", outer}: cord{32, 21},
			Vertex{"LF", inner}: cord{15, 28},
			Vertex{"QG", outer}: cord{32, 23},
			Vertex{"QG", inner}: cord{26, 17},
			Vertex{"VT", outer}: cord{32, 11},
			Vertex{"VT", inner}: cord{26, 23},
			Vertex{"YN", outer}: cord{2, 23},
			Vertex{"YN", inner}: cord{26, 13},
			Vertex{"ZZ", outer}: cord{2, 17},
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
		if g.ShortestPath(Vertex{"AA", outer}, Vertex{"ZZ", outer}) != tc.shortestPath {
			t.Errorf("g.ShortestPath(AA, ZZ) = %d, want %d", g.ShortestPath(Vertex{"AA", outer}, Vertex{"ZZ", outer}), tc.shortestPath)
		}
	}
}
