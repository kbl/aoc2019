package main

import (
	"aoc"
	"aoc/intcode"
	"fmt"
)

var debug = false

type cord struct {
	x, y int
}

type orientation int
type turn int
type color int

const (
	black = 0
	white = 1
)

const (
	turnLeft  = 0
	turnRight = 1
)

const (
	up    = 0
	right = 1
	down  = 2
	left  = 3
)

var moves = map[orientation][]int{
	up:    []int{0, 1},
	down:  []int{0, -1},
	left:  []int{-1, 0},
	right: []int{1, 0},
}

var turns = map[orientation]map[turn]orientation{
	up: map[turn]orientation{
		turnLeft:  left,
		turnRight: right,
	},
	right: map[turn]orientation{
		turnLeft:  up,
		turnRight: down,
	},
	down: map[turn]orientation{
		turnLeft:  right,
		turnRight: left,
	},
	left: map[turn]orientation{
		turnLeft:  down,
		turnRight: up,
	},
}

type Robot struct {
	c      cord
	o      orientation
	canvas map[cord]color
	ic     *intcode.Intcode
	done   bool
}

func NewRobot(ic *intcode.Intcode) *Robot {
	return &Robot{
		canvas: map[cord]color{},
		ic:     ic,
		done:   false,
	}
}

func (r *Robot) Step() {
	if r.done {
		panic("I'm done!")
	}
	var cordColor color = black
	if c, ok := r.canvas[r.c]; ok {
		cordColor = c
	}
	r.ic.AddInput(int(cordColor))

	toColor, exitMode := r.ic.Output()
	if exitMode == intcode.HaltMode {
		r.done = true
		return
	}
	toTurn, exitMode := r.ic.Output()
	if exitMode == intcode.HaltMode {
		r.done = true
		return
	}

	if toColor != black && toColor != white {
		panic(fmt.Sprintf("Illegal color %d!", toColor))
	}
	if toTurn != turnLeft && toTurn != turnRight {
		panic(fmt.Sprintf("Illegal turn %d!", toTurn))
	}

	r.canvas[r.c] = color(toColor)
	r.o = turns[r.o][turn(toTurn)]
	r.c = cord{r.c.x + moves[r.o][0], r.c.y + moves[r.o][1]}
}

func (r *Robot) Paint() {
	for !r.done {
		r.Step()
	}
}

func main() {
	inputFilePath := aoc.InputArg()
	Main(inputFilePath)
}

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)

	ic := intcode.NewIntcode(intcode.NewMemory(lines[0]))
	r := NewRobot(ic)
	r.Paint()
	fmt.Printf("Exercise 1: %d\n", len(r.canvas))

	ic = intcode.NewIntcode(intcode.NewMemory(lines[0]))
	r = NewRobot(ic)
	r.canvas[cord{0, 0}] = white
	r.Paint()
	fmt.Println("Exercise 2:")
	printCanvas(r.canvas)
}

func printCanvas(canvas map[cord]color) {
	minx := 0
	maxx := 0
	miny := 0
	maxy := 0

	for c := range canvas {
		if c.x > maxx {
			maxx = c.x
		}
		if c.x < minx {
			minx = c.x
		}
		if c.y > maxy {
			maxy = c.y
		}
		if c.y < miny {
			miny = c.y
		}
	}

	sColor := map[color]string{
		white: "▮",
		black: " ",
	}

	for y := maxy + 1; y >= miny-1; y-- {
		for x := minx - 1; x <= maxx+1; x++ {
			var toPaint color = black
			if c, ok := canvas[cord{x, y}]; ok {
				toPaint = c
			}
			fmt.Print(sColor[toPaint])
		}
		fmt.Println()
	}
}
