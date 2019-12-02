package main

import (
	"aoc"
	"fmt"
)

func main() {
	inputFilePath := aoc.InputArg()
	Main(inputFilePath)
}

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)
	fmt.Printf("Read %d lines!\n", len(lines))
}

func Program(input []int) int {
	index := 0
	for {
		operator := input[index]
		if operator == 99 {
			break
		}
		if operator == 1 {
			arg1Index := input[index+1]
			arg2Index := input[index+2]
			outputIndex := input[index+3]
			input[outputIndex] = add(input[arg1Index], input[arg2Index])
			index += 4
		} else if operator == 2 {
			arg1Index := input[index+1]
			arg2Index := input[index+2]
			outputIndex := input[index+3]
			input[outputIndex] = multiply(input[arg1Index], input[arg2Index])
			index += 4
		} else {
			index += 1
		}
	}
	return input[0]
}

func add(arg1, arg2 int) int {
	return arg1 + arg2
}

func multiply(arg1, arg2 int) int {
	return arg1 * arg2
}
