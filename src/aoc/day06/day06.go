package main

import (
	"aoc"
	"fmt"
	"log"
	"strings"
)

func main() {
	inputFilePath := aoc.InputArg()
	Main(inputFilePath)
}

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)
	orbits := NewOrbits(lines)
	fmt.Printf("Exercise 1: %d\n", orbits.AllOrbitsCount())
	fmt.Printf("Exercise 2: %d\n", orbits.OrbitalTransfers("YOU", "SAN"))
}

type Orbits struct {
	o  map[string][]string
	oc map[string]int
}

const com = "COM"

func NewOrbits(orbits []string) *Orbits {
	parsedOrbits := map[string][]string{}
	for _, l := range orbits {
		planets := strings.Split(l, ")")
		p0, p1 := planets[0], planets[1]
		if v, ok := parsedOrbits[p0]; ok {
			parsedOrbits[p0] = append(v, p1)
		} else {
			parsedOrbits[p0] = []string{p1}
		}
	}

	orbitCount := map[string]int{com: 0}
	planets := []string{com}

	for len(planets) > 0 {
		planet := planets[0]
		planets = planets[1:]

		for _, p := range parsedOrbits[planet] {
			if _, ok := orbitCount[p]; ok {
				log.Fatal(p)
			}
			orbitCount[p] = orbitCount[planet] + 1
			planets = append(planets, p)
		}
	}

	return &Orbits{parsedOrbits, orbitCount}
}

func (o *Orbits) commonAncestor(p1, p2 string) string {
	ancestors := map[string][]string{com: {}}
	planets := []string{com}

	for len(planets) > 0 {
		planet := planets[0]
		planets = planets[1:]

		for _, p := range o.o[planet] {
			ancestors[p] = append(ancestors[planet], planet)
			planets = append(planets, p)
		}
	}

	p1a := ancestors[p1]
	p2a := ancestors[p2]
	commonOrbit := ""

	for i, p := range p1a {
		if p2a[i] == p {
			commonOrbit = p
		} else {
			break
		}
	}

	return commonOrbit
}

func (o *Orbits) AllOrbitsCount() int {
	allOrbits := 0
	for _, count := range o.oc {
		allOrbits += count
	}
	return allOrbits
}

func (o *Orbits) OrbitalTransfers(p1, p2 string) int {
	commonAncestor := o.commonAncestor(p1, p2)
	oc := o.oc[commonAncestor]
	return o.oc[p1] - oc + o.oc[p2] - oc - 2
}
