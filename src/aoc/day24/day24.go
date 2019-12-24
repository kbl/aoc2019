package main

import (
	"aoc"
	"fmt"
	"math/bits"
	"strings"
)

const (
	bug    = '#'
	empty  = '.'
	maxlen = 5
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
		seen[grid] = true
		grid = evolve(grid)
	}
	return int(grid)
}

func bit(x, y int) int {
	return 1 << (y*maxlen + x)
}

func str(grid int) string {
	repr := []string{}
	for y := 0; y < maxlen; y++ {
		row := []rune{}
		for x := 0; x < maxlen; x++ {
			tile := empty
			if grid&bit(x, y) > 0 {
				tile = bug
			}
			row = append(row, tile)
		}
		repr = append(repr, string(row))
	}
	return strings.Join(repr, "\n")
}

func isBug(grid int, x, y int) bool {
	return grid&bit(x, y) > 0
}

func adjacent(x, y int) int {
	adj := 0
	if x > 0 {
		adj |= bit(x-1, y)
	}
	if y > 0 {
		adj |= bit(x, y-1)
	}
	if x < maxlen-1 {
		adj |= bit(x+1, y)
	}
	if y < maxlen-1 {
		adj |= bit(x, y+1)
	}
	return adj
}

func evolve(grid int) int {
	next := 0
	for y := 0; y < maxlen; y++ {
		for x := 0; x < maxlen; x++ {
			hm := bits.OnesCount(uint(grid & adjacent(x, y)))
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

func parse(lines []string) int {
	var grid int = 0
	for y, l := range lines {
		for x, r := range l {
			if r == bug {
				grid |= bit(x, y)
			}
		}
	}
	return grid
}
