package main

import (
	"aoc"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func (p *point) distance() int {
	return int(math.Abs(float64(p.x)) + math.Abs(float64(p.y)))
}

var moves = map[byte][]int{
	'L': {-1, 0},
	'R': {1, 0},
	'D': {0, -1},
	'U': {0, 1},
}

func (p *point) move(direction byte) *point {
	return &point{
		p.x + moves[direction][0],
		p.y + moves[direction][1],
	}
}

func (p *point) trail(instruction string) *[]point {
	distance, err := strconv.Atoi(instruction[1:])
	if err != nil {
		log.Fatal(err)
	}
	var trail []point
	for i := 1; i <= distance; i++ {
		p = p.move(instruction[0])
		trail = append(trail, *p)
	}
	return &trail
}

type Wires struct {
	wires []map[point]int
}

func (w *Wires) Wire(instructions []string) {
	wire := make(map[point]int)
	w.wires = append(w.wires, wire)

	currentPosition := &point{0, 0}
	length := 0
	for _, instruction := range instructions {
		for _, p := range *currentPosition.trail(instruction) {
			length += 1
			wire[p] = length
			currentPosition = &p
		}
	}
}

func (w *Wires) ClosestDistance() int {
	closest := -1
	w0 := w.wires[0]
	w1 := w.wires[1]

	for p := range w0 {
		if _, ok := w1[p]; !ok {
			continue
		}

		if closest == -1 {
			closest = p.distance()
		} else if p.distance() < closest {
			closest = p.distance()
		}
	}
	return closest
}

func (w *Wires) ShortestPath() int {
	closest := -1
	w0 := w.wires[0]
	w1 := w.wires[1]

	for p := range w0 {
		if _, ok := w1[p]; !ok {
			continue
		}

		steps := w1[p] + w0[p]
		if closest == -1 || closest > steps {
			closest = steps
		}
	}

	return closest
}

func main() {
	inputFilePath := aoc.InputArg()
	Main(inputFilePath)
}

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)
	directions := parseInput(lines)

	wires := Wires{}
	wires.Wire(directions[0])
	wires.Wire(directions[1])

	fmt.Printf("Exercise 1: %d\n", wires.ClosestDistance())
	fmt.Printf("Exercise 2: %d\n", wires.ShortestPath())
}

func parseInput(lines []string) [][]string {
	return [][]string{
		strings.Split(lines[0], ","),
		strings.Split(lines[1], ","),
	}
}
