package main

import (
	"aoc"
	"aoc/intcode"
	"fmt"
)

func main() {
	inputFilePath := aoc.InputArg()
	Main(inputFilePath)
}

type cord struct {
	x, y int
}

type mode int

const (
	stationary = 0
	pulled     = 1
)

func Main(inputFilePath string) {
	memory := aoc.Read(inputFilePath)[0]
	fmt.Println("Exercise 1:", exercise1(memory))
	fmt.Println("Exercise 2:", exercise2(memory))
}

func exercise1(memory string) int {
	how_many := 0

	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			switch check(x, y, memory) {
			case stationary:
				continue
			case pulled:
				how_many++
			default:
				panic("Unknown output!")
			}
		}
	}

	return how_many
}

func exercise2(memory string) int {
	// 2nd row is empty ;)
	x, y := 5, 2

	for {
		for {
			output := check(x, y, memory)
			if output == pulled {
				break
			} else {
				x++
			}
		}
		for localX := x; check(localX+99, y, memory) == pulled; localX++ {
			if check(localX, y+99, memory) == pulled && check(localX+99, y+99, memory) == pulled {
				return localX*10000 + y
			}
		}
		y++
	}
}

func check(x, y int, memory string) mode {
	cpu := intcode.NewIntcode(intcode.NewMemory(memory))
	cpu.AddInput(x)
	cpu.AddInput(y)
	o, m := cpu.Output()

	if m != intcode.OutputMode {
		panic("Illegal output mode!")
	}

	return mode(o)
}
