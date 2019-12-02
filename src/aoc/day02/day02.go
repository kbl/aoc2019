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
	strNumbers := strings.Split(lines[0], ",")
	var memory []int
	for _, sn := range strNumbers {
		m, err := strconv.Atoi(sn)
		if err != nil {
			log.Fatal(err)
		}
		memory = append(memory, m)
	}

	fmt.Printf("Exercise 1:")
	exercise1(memory)
	fmt.Printf("Exercise 2:")
	exercise2(memory)
}

func exercise1(memory []int) {
	memCopy := make([]int, len(memory))
	copy(memCopy, memory[:])
	memCopy[1] = 12
	memCopy[2] = 2
	fmt.Println(Intcode(memCopy))
}

func exercise2(memory []int) {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			memCopy := make([]int, len(memory))
			copy(memCopy, memory[:])
			memCopy[1] = noun
			memCopy[2] = verb
			fmt.Println(noun)
			if Intcode(memCopy) == 19690720 {
				fmt.Println(100*noun + verb)
				return
			}
		}
	}
}

func Intcode(memory []int) int {
	index := 0
	for {
		operator := memory[index]
		if operator == 99 {
			break
		}
		if operator == 1 {
			arg1Index := memory[index+1]
			arg2Index := memory[index+2]
			outputIndex := memory[index+3]
			memory[outputIndex] = add(memory[arg1Index], memory[arg2Index])
			index += 4
		} else if operator == 2 {
			arg1Index := memory[index+1]
			arg2Index := memory[index+2]
			outputIndex := memory[index+3]
			memory[outputIndex] = multiply(memory[arg1Index], memory[arg2Index])
			index += 4
		} else {
			index += 1
		}
	}
	return memory[0]
}

func add(arg1, arg2 int) int {
	return arg1 + arg2
}

func multiply(arg1, arg2 int) int {
	return arg1 * arg2
}
