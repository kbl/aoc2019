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

const (
	stationary = 0
	pulled     = 1
)

func Main(inputFilePath string) {
	memory := aoc.Read(inputFilePath)[0]
	fmt.Println("Exercise 1:", exercise1(memory))
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

func check(x, y int, memory string) int {
	cpu := intcode.NewIntcode(intcode.NewMemory(memory))
	cpu.AddInput(x)
	cpu.AddInput(y)
	o, m := cpu.Output()

	if m != intcode.OutputMode {
		fmt.Println(m, intcode.HaltMode, intcode.OutputMode)
		panic("AAA")
	}
	return o
}
