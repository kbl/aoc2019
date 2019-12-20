package main

import (
	"reflect"
	"testing"
)

type testCase struct {
	maze                  string
	vertices, edges       int
	entrances             map[Vertex]cord
	shortestPath          int
	recursiveShortestPath int
}

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
		shortestPath:          23,
		recursiveShortestPath: 27,
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
		shortestPath:          58,
		recursiveShortestPath: -1,
	},
	testCase{
		maze: `             Z L X W       C                 
             Z P Q B       K                 
  ###########.#.#.#.#######.###############  
  #...#.......#.#.......#.#.......#.#.#...#  
  ###.#.#.#.#.#.#.#.###.#.#.#######.#.#.###  
  #.#...#.#.#...#.#.#...#...#...#.#.......#  
  #.###.#######.###.###.#.###.###.#.#######  
  #...#.......#.#...#...#.............#...#  
  #.#########.#######.#.#######.#######.###  
  #...#.#    F       R I       Z    #.#.#.#  
  #.###.#    D       E C       H    #.#.#.#  
  #.#...#                           #...#.#  
  #.###.#                           #.###.#  
  #.#....OA                       WB..#.#..ZH
  #.###.#                           #.#.#.#  
CJ......#                           #.....#  
  #######                           #######  
  #.#....CK                         #......IC
  #.###.#                           #.###.#  
  #.....#                           #...#.#  
  ###.###                           #.#.#.#  
XF....#.#                         RF..#.#.#  
  #####.#                           #######  
  #......CJ                       NM..#...#  
  ###.#.#                           #.###.#  
RE....#.#                           #......RF
  ###.###        X   X       L      #.#.#.#  
  #.....#        F   Q       P      #.#.#.#  
  ###.###########.###.#######.#########.###  
  #.....#...#.....#.......#...#.....#.#...#  
  #####.#.###.#######.#######.###.###.#.#.#  
  #.......#.......#.#.#.#.#...#...#...#.#.#  
  #####.###.#####.#.#.#.#.###.###.#.###.###  
  #.......#.....#.#...#...............#...#  
  #############.#.#.###.###################  
               A O F   N                     
               A A D   M                     `,
		vertices: 15,
		edges:    21,
		entrances: map[Vertex]cord{
			Vertex{"AA", outer}: cord{15, 34},
			Vertex{"CJ", outer}: cord{2, 15},
			Vertex{"CJ", inner}: cord{8, 23},
			Vertex{"CK", outer}: cord{27, 2},
			Vertex{"CK", inner}: cord{8, 17},
			Vertex{"FD", outer}: cord{19, 34},
			Vertex{"FD", inner}: cord{13, 8},
			Vertex{"IC", outer}: cord{42, 17},
			Vertex{"IC", inner}: cord{23, 8},
			Vertex{"LP", outer}: cord{15, 2},
			Vertex{"LP", inner}: cord{29, 28},
			Vertex{"NM", outer}: cord{23, 34},
			Vertex{"NM", inner}: cord{36, 23},
			Vertex{"OA", outer}: cord{17, 34},
			Vertex{"OA", inner}: cord{8, 13},
			Vertex{"RE", outer}: cord{2, 25},
			Vertex{"RE", inner}: cord{21, 8},
			Vertex{"RF", outer}: cord{42, 25},
			Vertex{"RF", inner}: cord{36, 21},
			Vertex{"WB", outer}: cord{19, 2},
			Vertex{"WB", inner}: cord{36, 13},
			Vertex{"XF", outer}: cord{2, 21},
			Vertex{"XF", inner}: cord{17, 28},
			Vertex{"XQ", outer}: cord{17, 2},
			Vertex{"XQ", inner}: cord{21, 28},
			Vertex{"ZH", outer}: cord{42, 13},
			Vertex{"ZH", inner}: cord{31, 8},
			Vertex{"ZZ", outer}: cord{13, 2},
		},
		shortestPath:          77,
		recursiveShortestPath: 396,
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
		if tc.recursiveShortestPath != -1 {
			if g.RecursiveShortestPath(Vertex{"AA", outer}, Vertex{"ZZ", outer}) != tc.recursiveShortestPath {
				t.Errorf("g.RecursiveShortestPath(AA, ZZ) = %d, want %d", g.RecursiveShortestPath(Vertex{"AA", outer}, Vertex{"ZZ", outer}), tc.recursiveShortestPath)
			}
		}
	}
}
