package main

import (
	"aoc"
	"aoc/intcode"
	"fmt"
)

var debug = true

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

func (r *Robot) Turn(t turn) {
}

func NewRobot(ic *intcode.Intcode) *Robot {
	return &Robot{
		canvas: map[cord]color{},
		ic:     ic,
		done:   false,
	}
}

// 7176 too high
// 7175 too high

func (r *Robot) Step() {
	if r.done {
		fmt.Println(len(r.canvas))
		panic("I'm done!")
	}
	var cordColor color = black
	if v, ok := r.canvas[r.c]; ok {
		cordColor = v
	}
	r.ic.AddInput(int(cordColor))

	toColor, exitMode := r.ic.Output()
	if exitMode == intcode.HaltMode {
		fmt.Println("Halting 1!")
		r.done = true
		return
	}
	toTurn, exitMode := r.ic.Output()
	if exitMode == intcode.HaltMode {
		fmt.Println("Halting 2!")
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

	newOrientation := turns[r.o][turn(toTurn)]
	if debug {
		fmt.Printf("I'm oriented %v, turning %v to %v.\n", sOrientation[r.o], sTurn[turn(toTurn)], sOrientation[newOrientation])
	}
	r.o = newOrientation

	m := moves[r.o]
	newCord := cord{r.c.x + m[0], r.c.y + m[1]}
	if debug {
		fmt.Printf("Being oriented %v I'm moving from %v to %v.\n", sOrientation[r.o], r.c, newCord)
	}
	r.c = newCord
}

var sTurn = map[turn]string{
	turnLeft:  " left",
	turnRight: "right",
}

var sOrientation = map[orientation]string{
	up:    "   up",
	down:  " down",
	left:  " left",
	right: "right",
}

func main() {
	inputFilePath := aoc.InputArg()
	Main(inputFilePath)
}

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)
	intcode := intcode.NewIntcode(intcode.NewMemory(lines[0]))
	r := NewRobot(intcode)
	for {
		r.Step()
	}
	fmt.Println(r)
}
