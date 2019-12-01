package main

import (
	"aoc"
	"fmt"
	"log"
	"strconv"
)

func main() {
	lines := aoc.ReadFromArgs()
	solve(lines, fuel)
	solve(lines, totalFuel)
}

func solve(input []string, fuelCalc func(int) int) {
	sum := 0
	for _, l := range input {
		n, err := strconv.Atoi(l)
		if err != nil {
			log.Fatal(err)
		}
		sum += fuelCalc(n)
	}
	fmt.Println(sum)
}

func fuel(mass int) int {
	f := mass/3 - 2
	if f < 0 {
		return 0
	}
	return f
}

func totalFuel(mass int) int {
	f := 0
	fuelMass := mass
	for fuelMass > 0 {
		fuelMass = fuel(fuelMass)
		f += fuelMass
	}
	return f
}
