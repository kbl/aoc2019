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
	fmt.Println(str(grid))
	fmt.Println("Exercise 1:", exercise1(grid))
	fmt.Println(str(grid))
	fmt.Println("Exercise 2:", exercise2(10, grid))
}

func exercise1(grid int) int {
	seen := map[int]bool{}
	for !seen[grid] {
		seen[grid] = true
		grid = evolve(grid)
	}
	return int(grid)
}

func exercise2(steps, grid int) int {
	grids := map[int]int{0: grid}
	for s := 1; s <= steps; s++ {
		grids = evolveRecursive(grids)
		minLevel, maxLevel := 0, 0

		for l, _ := range grids {
			if l > maxLevel {
				maxLevel = l
			}
			if l < minLevel {
				minLevel = l
			}
		}
		fmt.Println(minLevel, maxLevel)

		fmt.Println("STEP", s)
		for l := minLevel; l <= maxLevel; l++ {
			fmt.Println("Level", l)
			fmt.Println(str(grids[l]))
			fmt.Println()
		}
		fmt.Println()
	}
	// 674 too low

	hm := 0
	for _, g := range grids {
		hm += bits.OnesCount(uint(g))
	}
	return hm
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

func adjacentLevel(level, x, y int) int {
	adj := 0

	if level == 1 {
		if x == 2 && y == 1 {
			for i := 0; i < maxlen; i++ {
				adj |= bit(i, 0)
			}
		}
		if x == 2 && y == 3 {
			for i := 0; i < maxlen; i++ {
				adj |= bit(i, 4)
			}
		}

		if x == 1 && y == 2 {
			for i := 0; i < maxlen; i++ {
				adj |= bit(0, i)
			}
		}
		if x == 3 && y == 2 {
			for i := 0; i < maxlen; i++ {
				adj |= bit(4, i)
			}
		}
		return adj
	}

	if x == 0 {
		adj |= bit(1, 2)
	}
	if y == 0 {
		adj |= bit(2, 1)
	}
	if x == maxlen-1 {
		adj |= bit(3, 2)
	}
	if y == maxlen-1 {
		adj |= bit(2, 3)
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

func evolveRecursive(grids map[int]int) map[int]int {
	next := map[int]int{}
	level := 0

	for {
		nextGrid := 0

		for y := 0; y < maxlen; y++ {
			for x := 0; x < maxlen; x++ {
				if x == 2 && y == 2 {
					continue
				}
				hm := bits.OnesCount(uint(grids[level] & adjacent(x, y)))
				hm += bits.OnesCount(uint(grids[level-1] & adjacentLevel(-1, x, y)))
				hm += bits.OnesCount(uint(grids[level+1] & adjacentLevel(+1, x, y)))
				if isBug(grids[level], x, y) && hm == 1 {
					nextGrid |= bit(x, y)
				}
				if !isBug(grids[level], x, y) && (hm == 1 || hm == 2) {
					nextGrid |= bit(x, y)
				}
			}
		}

		if nextGrid == 0 {
			break
		}

		next[level] = nextGrid
		level--
	}

	level = 1
	for {
		nextGrid := 0

		for y := 0; y < maxlen; y++ {
			for x := 0; x < maxlen; x++ {
				if x == 2 && y == 2 {
					continue
				}
				hm := bits.OnesCount(uint(grids[level] & adjacent(x, y)))
				hm += bits.OnesCount(uint(grids[level-1] & adjacentLevel(-1, x, y)))
				hm += bits.OnesCount(uint(grids[level+1] & adjacentLevel(1, x, y)))
				if isBug(grids[level], x, y) && hm == 1 {
					nextGrid |= bit(x, y)
				}
				if !isBug(grids[level], x, y) && (hm == 1 || hm == 2) {
					nextGrid |= bit(x, y)
				}
			}
		}

		if nextGrid == 0 {
			break
		}

		next[level] = nextGrid
		level++
	}

	return next
}

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
