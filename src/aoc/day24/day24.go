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
	fmt.Println(first(grid))
}

func first(grid int) int {
	seen := map[int]bool{}
	for !seen[grid] {
		fmt.Println(str(grid))
		fmt.Println()
		seen[grid] = true
		grid = evolve(grid)
	}
	return int(grid)
}

func bit(x, y int) int {
	return 1 << (y*5 + x)
}

func str(grid int) string {
	repr := []string{}
	for y := 0; y < 5; y++ {
		row := []rune{}
		for x := 0; x < 5; x++ {
			if grid&bit(x, y) > 0 {
				row = append(row, '#')
			} else {
				row = append(row, '.')
			}
		}
		repr = append(repr, string(row))
	}
	return strings.Join(repr, "\n")
}

func isBug(grid int, x, y int) bool {
	return grid&bit(x, y) > 0
}

func evolve(grid int) int {
	next := 0
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			hm := 0
			if x-1 >= 0 && isBug(grid, x-1, y) {
				hm++
			}
			if x+1 < 5 && isBug(grid, x+1, y) {
				hm++
			}
			if y-1 >= 0 && isBug(grid, x, y-1) {
				hm++
			}
			if y+1 < 5 && isBug(grid, x, y+1) {
				hm++
			}
			if isBug(grid, x, y) && hm == 1 {
				next |= bit(x, y)
			}
			if !isBug(grid, x, y) && (hm == 1 || hm == 2) {
				next |= bit(x, y)
			}
		}
	}
	return next
}

// 00 10 20 30 40
// 01 11 21 31 41
// 02 12 ?? 32 42
// 03 13 23 33 43
// 04 14 24 34 44
//
// xy
//
// y = 0
//   look += [level - 1][1][2]
// y = 4
//   look += [level - 1][3][2]
// x = 0
//   look += [level - 1][2][1]
// x = 4
//   look += [level - 1][2][3]
// x == 2 && y == 1
//   look += [level + 1][0][0] += [level + 1][0][0]
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

func parse(lines []string) int {
	var grid int = 0
	for y, l := range lines {
		for x, r := range l {
			if r == '#' {
				grid |= bit(x, y)
			}
		}
	}
	return grid
}
