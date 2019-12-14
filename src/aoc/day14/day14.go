package main

import (
	"aoc"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	inputFilePath := aoc.InputArg()
	Main(inputFilePath)
}

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)
	reactions := NewReactions(lines)

	fmt.Println("Exercise 1:", reactions.Process())
}

func (r Reactions) Process() int {
	required := map[string]int{}
	for chemical, quantity := range r.r["FUEL"].from {
		required[chemical] = quantity
	}

	anythingRequired := true
	for anythingRequired {
		anythingRequired = false
		for required_chemical, required_quantity := range required {
			if required_chemical == "ORE" || required_quantity <= 0 {
				continue
			}
			if r.r[required_chemical].quantity <= required_quantity {
				anythingRequired = true
				required[required_chemical] -= r.r[required_chemical].quantity
				for c, q := range r.r[required_chemical].from {
					if _, ok := required[c]; ok {
						required[c] += q
					} else {
						required[c] = q
					}
				}
			}
			if r.r[required_chemical].quantity > required_quantity {
				anythingRequired = true
				required[required_chemical] -= r.r[required_chemical].quantity
				for c, q := range r.r[required_chemical].from {
					if _, ok := required[c]; ok {
						required[c] += q
					} else {
						required[c] = q
					}
				}
			}
		}
	}

	fmt.Println(required)

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
