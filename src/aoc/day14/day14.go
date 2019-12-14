package main

import (
	"aoc"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	inputFilePath := aoc.InputArg()
	Main(inputFilePath)
}

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)
	reactions := NewReactions(lines)

	fmt.Println("Exercise 1:", reactions.OrePerFuel())
	fmt.Println("Exercise 2:", reactions.MaxFuelPerOre(1000000000000))
}

func (r Reactions) OrePerFuel() int {
	return r.Process(1)
}

func (r Reactions) MaxFuelPerOre(maxOre int) int {
	fuelMin := 0
	fuelMax := maxOre
	fuelIteration := maxOre / 2

	for fuelMin < fuelMax {
		time.Sleep(1 * time.Millisecond)
		iterationOre := r.Process(fuelIteration)
		if iterationOre == maxOre {
			return fuelIteration
		} else if iterationOre > maxOre {
			fuelMax = fuelIteration - 1
		} else {
			fuelMin = fuelIteration + 1
		}
		fuelIteration = fuelMin + (fuelMax-fuelMin)/2
	}
	return fuelIteration - 1
}

func (r Reactions) Process(amount int) int {
	required := map[string]int{}
	for chemical, quantity := range r.r["FUEL"].from {
		required[chemical] = quantity * amount
	}

	anythingRequired := true
	for anythingRequired {
		anythingRequired = false
		for required_chemical, required_quantity := range required {
			if required_chemical == "ORE" || required_quantity <= 0 {
				continue
			}
			produced_quantity := r.r[required_chemical].quantity
			multiplier := int(math.Ceil(float64(required_quantity) / float64(produced_quantity)))

			anythingRequired = true
			required[required_chemical] -= produced_quantity * multiplier
			for c, q := range r.r[required_chemical].from {
				required[c] += q * multiplier
			}
		}
	}

	return required["ORE"]
}

type Reaction struct {
	from     map[string]int
	to       string
	quantity int
}

func NewReaction(line string) Reaction {
	parts := strings.Split(line, " => ")
	input := strings.Split(parts[0], ", ")
	output := parts[1]

	parse := func(token string) (string, int) {
		tokens := strings.Split(token, " ")
		v, _ := strconv.Atoi(tokens[0])
		return tokens[1], v
	}

	to, quantity := parse(output)
	reaction := Reaction{map[string]int{}, to, quantity}

	for _, token := range input {
		c, v := parse(token)
		reaction.from[c] = v
	}

	return reaction
}

func (r Reaction) String() string {
	in := []string{}
	for c, v := range r.from {
		in = append(in, fmt.Sprintf("%d %s", v, c))
	}
	return fmt.Sprintf("%s => %d %s", strings.Join(in, ", "), r.quantity, r.to)
}

type Reactions struct {
	r map[string]Reaction
}

func NewReactions(lines []string) Reactions {
	reactions := map[string]Reaction{}
	for _, l := range lines {
		r := NewReaction(l)
		reactions[r.to] = r
	}
	return Reactions{reactions}
}
