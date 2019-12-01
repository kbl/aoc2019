package day01

import (
	"aoc"
	"fmt"
	"log"
	"strconv"
)

func Main() {
	lines := aoc.ReadFromArgs()
	solve(lines, Fuel)
	solve(lines, TotalFuel)
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

func Fuel(mass int) int {
	f := mass/3 - 2
	if f < 0 {
		return 0
	}
	return f
}

func TotalFuel(mass int) int {
	f := 0
	fuelMass := mass
	for fuelMass > 0 {
		fuelMass = Fuel(fuelMass)
		f += fuelMass
	}
	return f
}
