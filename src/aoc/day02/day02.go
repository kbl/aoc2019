package main

import (
	"aoc"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	inputFilePath := aoc.InputArg()
	Main(inputFilePath)
}

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)
	memory := parseMemory(lines[0])

	fmt.Printf("Exercise 1: ")
	exercise1(memory)
	fmt.Printf("Exercise 2: ")
	exercise2(memory)
}

func parseMemory(line string) *[]int {
	strNumbers := strings.Split(line, ",")
	var memory []int
	for _, sn := range strNumbers {
		m, err := strconv.Atoi(sn)
		if err != nil {
			log.Fatal(err)
		}
		memory = append(memory, m)
	}
	return &memory
}

func exercise1(memoryPtr *[]int) {
	memory := *memoryPtr
	memCopy := make([]int, len(memory))
	copy(memCopy, memory[:])
	memCopy[1] = 12
	memCopy[2] = 2
	fmt.Println(NewIntcode(&memCopy).Execute())
}

func exercise2(memoryPtr *[]int) {
	memory := *memoryPtr
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			memCopy := make([]int, len(memory))
			copy(memCopy, memory[:])
			memCopy[1] = noun
			memCopy[2] = verb
			if NewIntcode(&memCopy).Execute() == 19690720 {
				fmt.Println(100*noun + verb)
				return
			}
		}
	}
}

const (
	add      = 1
	multiply = 2
	halt     = 99
)

type Intcode struct {
	instructionPointer int
	memory             []int
}

func NewIntcode(memory *[]int) *Intcode {
	return &Intcode{
		0,
		*memory,
	}
}

func (i *Intcode) Execute() int {
	for {
		switch i.memory[i.instructionPointer] {
		case add:
			i.add()
		case multiply:
			i.multiply()
		case halt:
			i.halt()
			return i.memory[0]
		}
	}
}

func (i *Intcode) add() {
	param1Index := i.memory[i.instructionPointer+1]
	param2Index := i.memory[i.instructionPointer+2]
	outputIndex := i.memory[i.instructionPointer+3]
	i.memory[outputIndex] = i.memory[param1Index] + i.memory[param2Index]
	i.instructionPointer += 4
}

func (i *Intcode) multiply() {
	param1Index := i.memory[i.instructionPointer+1]
	param2Index := i.memory[i.instructionPointer+2]
	outputIndex := i.memory[i.instructionPointer+3]
	i.memory[outputIndex] = i.memory[param1Index] * i.memory[param2Index]
	i.instructionPointer += 4
}

func (i *Intcode) halt() {
	i.instructionPointer += 1
}
