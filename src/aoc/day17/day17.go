package main

import (
	"aoc"
	"aoc/intcode"
	"fmt"
	"strconv"
)

func main() {
	inputFilePath := aoc.InputArg()
	Main(inputFilePath)
}

type cord struct {
	x, y int
}

var moves = map[direction][]int{
	north: []int{0, -1},
	south: []int{0, 1},
	west:  []int{-1, 0},
	east:  []int{1, 0},
}

func (c cord) forward(d direction) cord {
	m := moves[d]
	return cord{
		x: c.x + m[0],
		y: c.y + m[1],
	}
}

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)
	fmt.Printf("Read %d lines!\n", len(lines))
	cpu := intcode.NewIntcode(intcode.NewMemory(lines[0]))
	grid := map[cord]string{}
	droid := cord{}
	x, y := 0, 0
	for {
		o, m := cpu.Output()

		if m == intcode.HaltMode {
			break
		}
		c := cord{x, y}
		switch o {
		case '#':
			grid[c] = "#"
			x++
		case '.':
			grid[c] = "."
			x++
		case '\n':
			x = 0
			y++
		case '>':
			droid = c
			grid[c] = ">"
			x++
		case '<':
			droid = c
			grid[c] = "<"
			x++
		case '^':
			droid = c
			grid[c] = "^"
			x++
		case 'v':
			droid = c
			grid[c] = "v"
			x++
		default:
			panic("Unknown character")
		}
	}

	fmt.Println("Exercise 1:", intersections(grid), droid)

	cpu = intcode.NewIntcode(intcode.NewMemory(lines[0]))
	route := findRoute(droid, grid)
	fmt.Println(route)

	cpu.SetMemory(0, 2)
}

[R 8 L 4 R 4 R 10 R 8 R 8 L 4 R 4 R 10 R 8 L 12 L 12 R 8 R 8 R 10 R 4 R 4 L 12 L 12 R 8 R 8 R 10 R 4 R 4 L 12 L 12 R 8 R 8 R 10 R 4 R 4 R 10 R 4 R 4 R 8 L 4 R 4 R 10 R]
[R 8 L 4 R 4 R 10 R 8 R 8 L 4 R 4 R 10 R 8 L 12 L 12 R 8 R 8 R 10 R 4 R 4 L 12 L 12 R 8 R 8 R 10 R 4 R 4 L 12 L 12 R 8 R 8 R 10 R 4 R 4 R 10 R 4 R 4 R 8 L 4 R 4 R 10 R]

const (
	north = "^"
	east  = ">"
	south = "v"
	west  = "<"
)

type direction string
type turn string

const (
	left  = "L"
	right = "R"
)

var turns = map[direction]map[turn]direction{
	north: map[turn]direction{
		left:  west,
		right: east,
	},
	east: map[turn]direction{
		left:  north,
		right: south,
	},
	south: map[turn]direction{
		left:  east,
		right: west,
	},
	west: map[turn]direction{
		left:  south,
		right: north,
	},
}

func findRoute(droid cord, grid map[cord]string) []string {
	// route := []string{}
	var d direction = north
	switch grid[droid] {
	case "^":
		d = north
	case ">":
		d = east
	case "v":
		d = south
	case "<":
		d = west
	default:
		panic("Wrong direction")
	}

	moved := 0
	sequence := []string{}
	for {
		newCord := droid.forward(d)
		if grid[newCord] == "#" {
			droid = newCord
			moved++
			continue
		}
		if grid[droid.forward(turns[d][left])] == "#" {
			d = turns[d][left]
			if moved > 0 {
				sequence = append(sequence, strconv.Itoa(moved))
				moved = 0
			}
			sequence = append(sequence, left)
		} else if grid[droid.forward(turns[d][right])] == "#" {
			d = turns[d][right]
			if moved > 0 {
				sequence = append(sequence, strconv.Itoa(moved))
				moved = 0
			}
			sequence = append(sequence, right)
		} else {
			return sequence
		}
	}
}

func intersections(grid map[cord]string) int {
	maxx, maxy := 0, 0
	for c, _ := range grid {
		if c.x > maxx {
			maxx = c.x
		}
		if c.y > maxy {
			maxy = c.y
		}
	}
	sum := 0
	for y := 0; y <= maxy; y++ {
		for x := 0; x <= maxx; x++ {
			tile := grid[cord{x, y}]
			if x == 0 || y == 0 || y == maxy || x == maxx {
				fmt.Print(tile)
				continue
			}
			if tile == "#" && grid[cord{x - 1, y}] == "#" && grid[cord{x + 1, y}] == "#" && grid[cord{x, y - 1}] == "#" && grid[cord{x, y + 1}] == "#" {
				sum += x * y
				fmt.Print("O")
			} else {
				fmt.Print(tile)
			}
		}
		fmt.Println()
	}
	return sum
}
