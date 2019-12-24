package main

import (
	"aoc"
	"fmt"
	"strings"
)

const (
	bug   = '#'
	empty = '.'
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

func first(grid [][]rune) int {
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

func str(grid [][]rune) string {
	repr := []string{}
	for _, row := range grid {
		repr = append(repr, string(row))
	}
	return strings.Join(repr, "\n")
}

func evolve(grid [][]rune) [][]rune {
	next := [][]rune{}
	for y, row := range grid {
		nextRow := []rune{}
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

func rating(grid [][]rune) int {
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

func parse(lines []string) [][]rune {
	grid := [][]rune{}
	for _, l := range lines {
		grid = append(grid, []rune(l))
	}
	return grid
}
