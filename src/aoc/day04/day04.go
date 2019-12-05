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
	tokens := strings.Split(lines[0], "-")
	min, err := strconv.Atoi(tokens[0])
	if err != nil {
		log.Fatal(err)
	}
	max, err := strconv.Atoi(tokens[1])
	if err != nil {
		log.Fatal(err)
	}

	exercise1Response := 0
	exercise2Response := 0
	for i := min; i <= max; i++ {
		if ValidPassword(i) {
			exercise1Response += 1
		}
		if ValidPassword2(i) {
			exercise2Response += 1
		}
	}
	fmt.Printf("Exercise 1: %d\n", exercise1Response)
	fmt.Printf("Exercise 2: %d\n", exercise2Response)
}

func ValidPassword(password int) bool {
	pstr := strconv.Itoa(password)
	previous := pstr[0]
	hadDuplicated := false
	for i := 1; i < len(pstr); i++ {
		if pstr[i] < previous {
			return false
		}
		if pstr[i] == previous {
			hadDuplicated = true
		}
		previous = pstr[i]
	}
	return hadDuplicated
}

func ValidPassword2(password int) bool {
	if !ValidPassword(password) {
		return false
	}
	pstr := strconv.Itoa(password)
	characters := map[rune]int{}
	for _, c := range pstr {
		if _, ok := characters[c]; ok {
			characters[c] += 1
		} else {
			characters[c] = 1
		}
	}

	for _, count := range characters {
		if count == 2 {
			return true
		}
	}
	return false
}
