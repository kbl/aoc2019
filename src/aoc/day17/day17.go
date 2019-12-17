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

	fmt.Println("Exercise 1:", intersections(grid))

	cpu = intcode.NewIntcode(intcode.NewMemory(lines[0]))
	route := findRoute(droid, grid)
	fmt.Println(route)
	// manual split
	// [R 8 L 4 R 4 R 10 R 8 R 8 L 4 R 4 R 10 R 8 L 12 L 12 R 8 R 8 R 10 R 4 R 4 L 12 L 12 R 8 R 8 R 10 R 4 R 4 L 12 L 12 R 8 R 8 R 10 R 4 R 4 R 10 R 4 R 4 R 8 L 4 R 4 R 10 R 8]
	// A =  R 8 L 4 R 4 R 10 R 8
	// [A                    A                    L 12 L 12 R 8 R 8 R 10 R 4 R 4 L 12 L 12 R 8 R 8 R 10 R 4 R 4 L 12 L 12 R 8 R 8 R 10 R 4 R 4 R 10 R 4 R 4 A                   ]
	// B = L 12 L 12 R 8 R 8
	// [A                    A                    B                 R 10 R 4 R 4 B                 R 10 R 4 R 4 B                 R 10 R 4 R 4 R 10 R 4 R 4 A                   ]
	// C = R 10 R 4 R 4
	// [A                    A                    B                 C            B                 C            B                 C            C            A                   ]
	cpu.SetMemory(0, 2)
	functions := map[string][]string{
		"A": []string{"R", "8", "L", "4", "R", "4", "R", "10", "R", "8"},
		"B": []string{"L", "12", "L", "12", "R", "8", "R", "8"},
		"C": []string{"R", "10", "R", "4", "R", "4"},
	}
	mainFunction := []string{"A", "A", "B", "C", "B", "C", "B", "C", "C", "A"}

	toAscii := func(str []string) []int {
		ascii := []int{}
		for i, s := range str {
			for _, c := range s {
				ascii = append(ascii, int(c))
			}
			if i < len(str)-1 {
				ascii = append(ascii, ',')
			}
		}
		return ascii
	}

	for _, v := range toAscii(mainFunction) {
		cpu.AddInput(v)
	}
	cpu.AddInput(int('\n'))

	for _, v := range toAscii(functions["A"]) {
		cpu.AddInput(v)
	}
	cpu.AddInput(int('\n'))

	for _, v := range toAscii(functions["B"]) {
		cpu.AddInput(v)
	}
	cpu.AddInput(int('\n'))

	for _, v := range toAscii(functions["C"]) {
		cpu.AddInput(v)
	}
	cpu.AddInput(int('\n'))

	cpu.AddInput(int('n'))
	cpu.AddInput(int('\n'))

	for {
		o, m := cpu.Output()
		if m == intcode.HaltMode {
			fmt.Println("Exercise 2:", o)
			break
		}
	}
}

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
			return append(sequence, strconv.Itoa(moved))
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
