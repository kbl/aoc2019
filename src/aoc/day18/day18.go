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
