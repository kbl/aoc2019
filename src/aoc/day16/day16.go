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

func AdvancedPhase(iterations int, input string) string {
	repeatFactor := 10000
	numberLength := repeatFactor * len(input)
	minimalOffset := numberLength / 2

	digits := parse(input)
	digitsStr := []string{}
	for _, d := range digits[:7] {
		digitsStr = append(digitsStr, strconv.Itoa(d))
	}
	offset, err := strconv.Atoi(strings.Join(digitsStr, ""))
	if err != nil {
		log.Fatal(err)
	}
	if offset < minimalOffset {
		panic(fmt.Sprintf("%d offset is smaller than %d!", offset, minimalOffset))
	}

	previous := make([]int, numberLength-offset)
	current := make([]int, numberLength-offset)

	for di := numberLength - 1; di >= offset; di-- {
		previous[di-offset] = digits[di%len(digits)]
	}
	current[len(current)-1] = previous[len(previous)-1]

	for i := 0; i < iterations; i++ {
		for di := len(current) - 2; di >= 0; di-- {
			current[di] = (current[di+1] + previous[di]) % 10
		}
		previous = current
	}

	resultStr := []string{}
	for i := 0; i < 8; i++ {
		resultStr = append(resultStr, strconv.Itoa(current[i]))
	}

	return strings.Join(resultStr, "")
}

func digitValueAfterIteration(digits []int, numberLength, iteration, index int) int {
	if iteration < 1 {
		panic(fmt.Sprintf("Illegal iteration value %d!", iteration))
	}

	previous := make([]int, numberLength-index)
	current := make([]int, numberLength-index)
	for di := numberLength - 1; di >= index; di-- {
		previous[di-index] = digits[di%len(digits)]
	}
	current[len(current)-1] = previous[len(previous)-1]

	for i := 0; i < iteration; i++ {
		for di := len(current) - 2; di >= 0; di-- {
			current[di] = (current[di+1] + previous[di]) % 10
		}
		previous = current
	}

	return current[0]
}

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)
	fmt.Printf("Exercise 1: %s\n", Phase(100, lines[0]))
	fmt.Printf("Exercise 2: %s\n", AdvancedPhase(100, lines[0]))
}
