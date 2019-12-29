package main

import (
	"aoc"
	"fmt"
	"math/bits"
	"strings"
)

const (
	bug    = '#'
	empty  = '.'
	maxlen = 5
)

func main() {
	inputFilePath := aoc.InputArg()
	Main(inputFilePath)
}

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)
	g := parse(lines)
	fmt.Println(g)
	fmt.Println("Exercise 1:", exercise1(g))
	fmt.Println(g)
	fmt.Println("Exercise 2:", exercise2(200, g))
}

func exercise1(g grid) int {
	seen := map[grid]bool{}
	for !seen[g] {
		seen[g] = true
		g = g.evolve()
	}
	return int(g)
}

func str(grids map[int]grid) {
	minLevel, maxLevel := 0, 0

	for l, _ := range grids {
		if l > maxLevel {
			maxLevel = l
		}
		if l < minLevel {
			minLevel = l
		}
	}
	fmt.Println(minLevel, maxLevel)

	for l := minLevel; l <= maxLevel; l++ {
		fmt.Println("Level", l)
		fmt.Println(grids[l])
		fmt.Println()
	}
	fmt.Println()
}

func exercise2(steps int, g grid) int {
	grids := map[int]grid{0: g}
	for s := 1; s <= steps; s++ {
		grids = evolveRecursive(grids)
	}
	// 674 too low
	str(grids)

	hm := 0
	for _, g := range grids {
		hm += bits.OnesCount(uint(g))
	}
	return hm
}

func bit(x, y int) int {
	return 1 << (y*maxlen + x)
}

type grid int

func (g grid) String() string {
	repr := []string{}
	for y := 0; y < maxlen; y++ {
		row := []rune{}
		for x := 0; x < maxlen; x++ {
			tile := empty
			if int(g)&bit(x, y) > 0 {
				tile = bug
			}
			row = append(row, tile)
		}
		repr = append(repr, string(row))
	}
	return strings.Join(repr, "\n")
}

func (g grid) isBug(x, y int) bool {
	return int(g)&bit(x, y) > 0
}

func adjacent(x, y int) int {
	adj := 0
	if x > 0 {
		adj |= bit(x-1, y)
	}
	if y > 0 {
		adj |= bit(x, y-1)
	}
	if x < maxlen-1 {
		adj |= bit(x+1, y)
	}
	if y < maxlen-1 {
		adj |= bit(x, y+1)
	}
	return adj
}

//           21
//
//     00 10 20 30 40
//     01 11 21 31 41
// 12  02 12 xy 32 42  32
//     03 13 23 33 43
//     04 14 24 34 44
//
//           23

func adjacentLevel(level, x, y int) int {
	adj := 0

	if level == 1 {
		if x == 2 && y == 1 {
			for i := 0; i < maxlen; i++ {
				adj |= bit(i, 0)
			}
		}
		if x == 2 && y == 3 {
			for i := 0; i < maxlen; i++ {
				adj |= bit(i, 4)
			}
		}

		if x == 1 && y == 2 {
			for i := 0; i < maxlen; i++ {
				adj |= bit(0, i)
			}
		}
		if x == 3 && y == 2 {
			for i := 0; i < maxlen; i++ {
				adj |= bit(4, i)
			}
		}
		return adj
	}

	if x == 0 {
		adj |= bit(1, 2)
	}
	if y == 0 {
		adj |= bit(2, 1)
	}
	if x == maxlen-1 {
		adj |= bit(3, 2)
	}
	if y == maxlen-1 {
		adj |= bit(2, 3)
	}

	return adj
}

func (g grid) evolve() grid {
	next := 0
	for y := 0; y < maxlen; y++ {
		for x := 0; x < maxlen; x++ {
			hm := bits.OnesCount(uint(int(g) & adjacent(x, y)))
			if g.isBug(x, y) && hm == 1 {
				next |= bit(x, y)
			}
			if !g.isBug(x, y) && (hm == 1 || hm == 2) {
				next |= bit(x, y)
			}
		}
	}
	return grid(next)
}

func evolveRecursive(grids map[int]grid) map[int]grid {
	next := map[int]grid{}
	level := 0

	for {
		nextGrid := 0

		for y := 0; y < maxlen; y++ {
			for x := 0; x < maxlen; x++ {
				if x == 2 && y == 2 {
					continue
				}
				hm := bits.OnesCount(uint(int(grids[level]) & adjacent(x, y)))
				hm += bits.OnesCount(uint(int(grids[level-1]) & adjacentLevel(-1, x, y)))
				hm += bits.OnesCount(uint(int(grids[level+1]) & adjacentLevel(+1, x, y)))
				if grids[level].isBug(x, y) && hm == 1 {
					nextGrid |= bit(x, y)
				}
				if !grids[level].isBug(x, y) && (hm == 1 || hm == 2) {
					nextGrid |= bit(x, y)
				}
			}
		}

		_, isDeeper := grids[level-1]
		if nextGrid == 0 && !isDeeper {
			break
		}

		next[level] = grid(nextGrid)
		level--
	}

	level = 1
	for {
		nextGrid := 0

		for y := 0; y < maxlen; y++ {
			for x := 0; x < maxlen; x++ {
				if x == 2 && y == 2 {
					continue
				}
				hm := bits.OnesCount(uint(int(grids[level]) & adjacent(x, y)))
				hm += bits.OnesCount(uint(int(grids[level-1]) & adjacentLevel(-1, x, y)))
				hm += bits.OnesCount(uint(int(grids[level+1]) & adjacentLevel(1, x, y)))
				if grids[level].isBug(x, y) && hm == 1 {
					nextGrid |= bit(x, y)
				}
				if !grids[level].isBug(x, y) && (hm == 1 || hm == 2) {
					nextGrid |= bit(x, y)
				}
			}
		}

		_, isDeeper := grids[level+1]
		if nextGrid == 0 && !isDeeper {
			break
		}

		next[level] = grid(nextGrid)
		level++
	}

	return next
}

func parse(lines []string) grid {
	var g = 0
	for y, l := range lines {
		for x, r := range l {
			if r == bug {
				g |= bit(x, y)
			}
		}
	}
	return grid(g)
}
