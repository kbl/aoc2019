package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		sum += totalFuel(n)
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
