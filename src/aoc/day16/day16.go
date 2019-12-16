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

func parse(line string) []int {
	digits := []int{}
	for _, d := range strings.Split(line, "") {
		n, err := strconv.Atoi(d)
		if err != nil {
			log.Fatal(err)
		}
		digits = append(digits, n)
	}
	return digits
}

// 0 1 0 -1
// 0 0 1 1 0 0 -1 -1
// 0 0 0 1 1 1 0 0 0 -1 -1 -1

func phaseMultiplier(phase, sequenceIndex int) int {
	basePhase := []int{0, 1, 0, -1}
	index := (sequenceIndex + 1) / phase
	return basePhase[index%len(basePhase)]
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func Phase(iterations int, input string) string {
	digits := parse(input)

	for i := 0; i < iterations; i++ {
		newDigits := []int{}
		for di := 0; di < len(digits); di++ {
			sum := 0
			for ddi, d := range digits {
				sum += (d * phaseMultiplier(di+1, ddi)) % 10
			}
			newDigits = append(newDigits, abs(sum%10))
		}
		digits = newDigits
	}

	dStr := []string{}
	for i := 0; i < 8; i++ {
		dStr = append(dStr, strconv.Itoa(digits[i]))
	}
	return strings.Join(dStr, "")
}

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)
	fmt.Printf("Exercise 1: %d\n", Phase(100, lines[0]))
}
