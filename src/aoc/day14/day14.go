package main

import (
	"aoc"
	"fmt"
	"log"
	"math"
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

// too high 580387

type Chemical string

type Reaction struct {
	in   map[Chemical]int
	out  Chemical
	outv int
}

func (r Reaction) String() string {
	strs := []string{}
	for c, q := range r.in {
		strs = append(strs, fmt.Sprintf("%d %s", q, c))
	}
	return fmt.Sprintf("%s => %d %s", strings.Join(strs, ", "), r.outv, r.out)
}

func (r Reaction) ReplaceInplace(r2 Reaction, onlyFull bool) bool {
	chemical := r2.out
	if _, ok := r.in[chemical]; !ok {
		return false
	}
	requiredQuantity := r.in[chemical]
	if onlyFull && requiredQuantity < r2.outv {
		return false
	}

	multiplier := requiredQuantity / r2.outv
	reminder := requiredQuantity % r2.outv

	if !onlyFull {
		multiplier = int(math.Ceil(float64(requiredQuantity) / float64(r2.outv)))
		reminder = 0
	}

	if reminder > 0 {
		r.in[chemical] = reminder
	} else {
		delete(r.in, chemical)
	}

	for c, q := range r2.in {
		if _, ok := r.in[c]; ok {
			r.in[c] += multiplier * q
		} else {
			r.in[c] = multiplier * q
		}
	}

	return true
}

func (r Reaction) clone() Reaction {
	or := Reaction{
		in:   map[Chemical]int{},
		out:  r.out,
		outv: r.outv,
	}
	for c, q := range r.in {
		or.in[c] = q
	}
	return or
}

func (r Reaction) Replace(r2 Reaction) Reaction {
	chemical := r2.out
	if _, ok := r.in[chemical]; !ok {
		return r
	}
	or := r.clone()

	multiplier := int(math.Ceil(float64(or.in[chemical]) / float64(r2.outv)))

	for c, q := range r2.in {
		if _, ok := or.in[c]; ok {
			or.in[c] += multiplier * q
		} else {
			or.in[c] = multiplier * q
		}
	}

	return or
}

func NewReaction(line string) Reaction {
	equation := strings.Split(line, " => ")

	processPart := func(equationPart string) map[Chemical]int {
		output := map[Chemical]int{}
		for _, chemical := range strings.Split(equationPart, ", ") {
			splits := strings.Split(chemical, " ")
			i, err := strconv.Atoi(splits[0])
			if err != nil {
				log.Fatal(err)
			}
			c := Chemical(splits[1])
			if _, ok := output[c]; ok {
				panic(fmt.Sprintf("%s already in inputs!", c))
			}
			output[c] = i
		}
		return output
	}

	r := Reaction{
		in: processPart(equation[0]),
	}

	for k, v := range processPart(equation[1]) {
		r.out = k
		r.outv = v
	}
	return r
}

type Reactions struct {
	r []Reaction
}

func NewReactions(lines []string) *Reactions {
	r := Reactions{}
	for _, l := range lines {
		r.r = append(r.r, NewReaction(l))
	}
	return &r
}

func (r *Reactions) String() string {
	repr := []string{}
	for _, reaction := range r.r {
		repr = append(repr, fmt.Sprint(reaction))
	}
	return strings.Join(repr, "\n")
}

func (r *Reactions) Process() int {
	r.Simplify()
	if len(r.r) != 1 {
		fmt.Println(r)
		panic("Too many reactions!")
	}
	fmt.Println(r.r)
	return r.r[0].in["ORE"]
}

func (r *Reactions) Simplify() {
	uberSimplified := true
	for uberSimplified {
		uberSimplified = false
		uberSimplified = uberSimplified || r.replaceFullOccurences()
		uberSimplified = uberSimplified || r.removeUnusedEquations()
		uberSimplified = uberSimplified || r.replaceSingleOccurences()
	}
}

func (r *Reactions) replaceFullOccurences() bool {
	replaced := false
	for i, r1 := range r.r {
		for j, r2 := range r.r {
			if i == j {
				continue
			}
			if r2.ReplaceInplace(r1, true) {
				replaced = true
			}
		}
	}
	return replaced
}

func (r *Reactions) removeUnusedEquations() bool {
	equationRemoved := false
	reactions := []Reaction{}

	for i, r1 := range r.r {
		chemicalUsed := false
		if r1.out == "FUEL" {
			reactions = append(reactions, r1)
			continue
		}
		for j, r2 := range r.r {
			if i == j {
				continue
			}
			if _, ok := r2.in[r1.out]; ok {
				reactions = append(reactions, r1)
				chemicalUsed = true
				break
			}
		}
		if !chemicalUsed {
			equationRemoved = true
		}
	}
	r.r = reactions

	return equationRemoved
}

func (r *Reactions) replaceSingleOccurences() bool {
	used := map[Chemical][]Reaction{}
	producing := map[Chemical]Reaction{}

	for _, r1 := range r.r {
		producing[r1.out] = r1
		for _, r2 := range r.r {
			if _, ok := r2.in[r1.out]; ok {
				used[r1.out] = append(used[r1.out], r2)
			}
		}
	}

	replaced := false

	for chemical, reactions := range used {
		if len(reactions) != 1 {
			continue
		}
		replaced = true
		output_used_once := producing[chemical]
		where_used := reactions[0]
		where_used.ReplaceInplace(output_used_once, false)
	}

	return replaced
}
