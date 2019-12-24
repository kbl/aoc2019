package main

import (
	"aoc"
	"fmt"
	"strings"
)

const (
	bug   = true
	empty = false
)

func main() {
	inputFilePath := aoc.InputArg()
	Main(inputFilePath)
}

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)
	grid := parse(lines)
	fmt.Printf("Read %d lines!\n", len(grid))
	fmt.Println(first(grid))
}

func first(grid [][]bool) int {
	seen := map[int]bool{}
	r := rating(grid)
	for !seen[r] {
		fmt.Println(str(grid))
		fmt.Println()
		seen[r] = true
		grid = evolve(grid)
		r = rating(grid)

	}
	fmt.Println(str(grid))
	fmt.Println()
	return r
}

func str(grid [][]bool) string {
	repr := []string{}
	for _, row := range grid {
		r := []rune{}
		for _, tile := range row {
			tileS := '.'
			if tile == bug {
				tileS = '#'
			}
			r = append(r, tileS)
		}
		repr = append(repr, string(r))
	}
	return strings.Join(repr, "\n")
}

func evolve(grid [][]bool) [][]bool {
	next := [][]bool{}
	for y, row := range grid {
		nextRow := []bool{}
		for x, tile := range row {
			hm := 0
			if x-1 >= 0 && grid[y][x-1] == bug {
				hm++
			}
			if x+1 < 5 && grid[y][x+1] == bug {
				hm++
			}
			if y-1 >= 0 && grid[y-1][x] == bug {
				hm++
			}
			if y+1 < 5 && grid[y+1][x] == bug {
				hm++
			}
			nextTile := empty
			if tile == bug && hm == 1 {
				nextTile = bug
			}
			if tile == empty && (hm == 1 || hm == 2) {
				nextTile = bug
			}
			nextRow = append(nextRow, nextTile)
		}
		next = append(next, nextRow)
	}
	return next
}

// 00 10 20 30 40
// 01 11 21 31 41
// 02 12 ?? 32 42
// 03 13 23 33 43
// 04 14 24 34 44
//
// y in [0, 4]
// look += [level + 1]
// 00
// 10
// 20
// 30
// 40
// 01
// 02
// 03
// 04

func evolveRecursive(grids map[int][][]bool) map[int][][]bool {
	level := 0
	nextGrids := map[int][][]bool{}
	for grid, ok := grids[level]; ok; {
		nextGrid := [][]bool{}
		for y, row := range grid {
			nextRow := []bool{}
			for x, tile := range row {
				if x == 2 && y == 2 {
					continue
				}

				hm := 0
				if x-1 >= 0 && grid[y][x-1] == bug {
					hm++
				}
				if x+1 < 5 && grid[y][x+1] == bug {
					hm++
				}
				if y-1 >= 0 && grid[y-1][x] == bug {
					hm++
				}
				if y+1 < 5 && grid[y+1][x] == bug {
					hm++
				}
				nextTile := empty
				if tile == bug && hm == 1 {
					nextTile = bug
				}
				if tile == empty && (hm == 1 || hm == 2) {
					nextTile = bug
				}
				nextRow = append(nextRow, nextTile)
			}
			nextGrid = append(nextGrid, nextRow)
		}
		nextGrids[level] = nextGrid
		level--
	}
	return nextGrids
}

func rating(grid [][]bool) int {
	rating := 0
	for y, row := range grid {
		for x, tile := range row {
			if tile == bug {
				rating += 1 << (y*5 + x)
			}
		}
	}
	return rating
}

func parse(lines []string) [][]bool {
	grid := [][]bool{}
	for _, l := range lines {
		row := []bool{}
		for _, r := range l {
			tileType := bug
			if r == '.' {
				tileType = empty
			}
			row = append(row, tileType)
		}
		grid = append(grid, row)
	}
	return grid
}
