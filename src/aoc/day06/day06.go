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
	orbits := map[string][]string{}
	for _, l := range lines {
		planets := strings.Split(l, ")")
		if v, ok := orbits[planets[0]]; ok {
			orbits[planets[0]] = append(v, planets[1])
		} else {
			orbits[planets[0]] = []string{planets[1]}
		}
	}
	orbitCount := map[string]int{
		"COM": 0,
	}
	planets := []string{"COM"}

	for len(planets) > 0 {
		planet := planets[0]
		planets = planets[1:]

		for _, p := range orbits[planet] {
			if _, ok := orbitCount[p]; ok {
				log.Fatal(p)
			}
			orbitCount[p] = orbitCount[planet] + 1
			planets = append(planets, p)
		}
	}

	allOrbits := 0
	for _, v := range orbitCount {
		allOrbits += v
	}

	fmt.Println(allOrbits)
	x := CommonAncestor(orbits, "SAN", "YOU")
	oc := orbitCount[x]
	fmt.Println(orbitCount["SAN"] - oc)
	fmt.Println(orbitCount["YOU"] - oc)
	fmt.Println(orbitCount["SAN"] - oc + orbitCount["YOU"] - oc - 2)
}

func CommonAncestor(orbits map[string][]string, p1, p2 string) string {
	ancestors := map[string][]string{"COM": {}}
	planets := []string{"COM"}

	for len(planets) > 0 {
		planet := planets[0]
		planets = planets[1:]

		for _, p := range orbits[planet] {
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
