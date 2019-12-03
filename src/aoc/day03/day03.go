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

func (p *point) trail(instruction string) *[]point {
	distance, err := strconv.Atoi(instruction[1:])
	if err != nil {
		log.Fatal(err)
	}
	var trail []point
	switch instruction[0] {
	case 'L':
		for i := 1; i <= distance; i++ {
			p := point{
				p.x - i,
				p.y,
			}
			trail = append(trail, p)
		}
	case 'R':
		for i := 1; i <= distance; i++ {
			p := point{
				p.x + i,
				p.y,
			}
			trail = append(trail, p)
		}
	case 'U':
		for i := 1; i <= distance; i++ {
			p := point{
				p.x,
				p.y + i,
			}
			trail = append(trail, p)
		}
	case 'D':
		for i := 1; i <= distance; i++ {
			p := point{
				p.x,
				p.y - i,
			}
			trail = append(trail, p)
		}
	}
	return &trail
}

type Wires struct {
	grid1, grid2 map[point]bool
	wire1, wire2 []string
}

func NewWires() *Wires {
	return &Wires{
		make(map[point]bool),
		make(map[point]bool),
		nil,
		nil,
	}
}

func (w *Wires) Wire1(wire []string) {
	w.wire1 = wire
	currentPosition := &point{0, 0}
	for _, instruction := range wire {
		trail := *currentPosition.trail(instruction)
		for _, p := range trail {
			w.grid1[p] = true
		}
		currentPosition = &trail[len(trail)-1]
	}
}

func (w *Wires) Wire2(wire []string) {
	w.wire2 = wire
	currentPosition := &point{0, 0}
	for _, instruction := range wire {
		trail := *currentPosition.trail(instruction)
		for _, p := range trail {
			w.grid2[p] = true
		}
		currentPosition = &trail[len(trail)-1]
	}
}

func (w *Wires) Closest() int {
	closest := -1
	for p, _ := range w.grid1 {
		if _, ok := w.grid2[p]; !ok {
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

func (w *Wires) Steps() int {
	intersections := make(map[point]int)
	for p, _ := range w.grid1 {
		if _, ok := w.grid2[p]; ok {
			intersections[p] = 0
		}
	}

	wire := w.wire1
	steps := 0
	currentPosition := &point{0, 0}
	for _, instruction := range wire {
		trail := *currentPosition.trail(instruction)
		for _, p := range trail {
			steps += 1
			if _, ok := intersections[p]; ok {
				intersections[p] += steps
			}
		}
		currentPosition = &trail[len(trail)-1]
	}

	wire = w.wire2
	steps = 0
	currentPosition = &point{0, 0}
	for _, instruction := range wire {
		trail := *currentPosition.trail(instruction)
		for _, p := range trail {
			steps += 1
			if _, ok := intersections[p]; ok {
				intersections[p] += steps
			}
		}
		currentPosition = &trail[len(trail)-1]
	}

	closest := -1
	for _, v := range intersections {
		if closest == -1 {
			closest = v
		} else if v < closest {
			closest = v
		}
	}

	return closest
}

func main() {
	inputFilePath := aoc.InputArg()
	lines := aoc.Read(inputFilePath)
	directions := parseInput(lines)
	Main(directions)
}

func parseInput(lines []string) [][]string {
	return [][]string{
		strings.Split(lines[0], ","),
		strings.Split(lines[1], ","),
	}
}

func Main(directions [][]string) {
	fmt.Printf("Read %s lines!\n", directions[0])
	fmt.Printf("Read %s lines!\n", directions[1])
	wires := NewWires()
	wires.Wire1(directions[0])
	wires.Wire2(directions[1])
}
