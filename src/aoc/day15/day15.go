package main

import (
	"aoc"
	"aoc/intcode"
	"fmt"
	"strings"
)

func main() {
	inputFilePath := aoc.InputArg()
	Main(inputFilePath)
}

const (
	north = 1
	south = 2
	west  = 3
	east  = 4
)

const (
	wall         = 0
	moved        = 1
	oxygenSystem = 2
)

var moveString = map[move]string{
	north: "^",
	south: "v",
	west:  "<",
	east:  ">",
}

var tileString = map[int]string{
	wall:         "#",
	oxygenSystem: "O",
	moved:        " ",
}

type move int

type Cord struct {
	x, y int
}

func (c Cord) Move(direction move) Cord {
	moves := map[move][2]int{
		north: {0, 1},
		south: {0, -1},
		west:  {-1, 0},
		east:  {1, 0},
	}
	return Cord{c.x + moves[direction][0], c.y + moves[direction][1]}
}

func (c Cord) Adjacent() map[move]Cord {
	return map[move]Cord{
		north: c.Move(north),
		south: c.Move(south),
		west:  c.Move(west),
		east:  c.Move(east),
	}
}

type Grid map[Cord]int

func (g Grid) String() string {
	minx := 0
	maxx := 0
	miny := 0
	maxy := 0
	for c, _ := range g {
		if c.x > maxx {
			maxx = c.x
		}
		if c.y > maxy {
			maxy = c.y
		}
		if c.x < minx {
			minx = c.x
		}
		if c.y < miny {
			miny = c.y
		}
	}

	grid := []string{}
	for y := maxy + 1; y >= miny-1; y-- {
		row := []string{}
		for x := minx - 1; x <= maxx+1; x++ {
			if x == 0 && y == 0 {
				row = append(row, "*")
			} else if c, ok := g[Cord{x, y}]; ok {
				row = append(row, tileString[c])
			} else {
				row = append(row, "?")
			}
		}
		grid = append(grid, strings.Join(row, ""))
	}
	return fmt.Sprintf(strings.Join(grid, "\n"))
}

type Droid struct {
	cpu      *intcode.Intcode
	grid     Grid
	position Cord
}

func NewDroid(cpu *intcode.Intcode) *Droid {
	return &Droid{
		cpu:  cpu,
		grid: Grid{Cord{0, 0}: moved},
	}
}

func (d *Droid) Discover() {
	toVisit := []Cord{}
	toVisitM := map[Cord]bool{}
	for _, c := range d.position.Adjacent() {
		toVisit = append(toVisit, c)
		toVisitM[c] = true
	}

	for len(toVisit) > 0 {
		next := toVisit[0]
		toVisit = toVisit[1:]
		delete(toVisitM, next)

		moveSequence := d.moveSequence(next)

		for _, moveCommand := range moveSequence {
			d.cpu.AddInput(int(moveCommand))
			tile, outputMode := d.cpu.Output()
			if outputMode != intcode.OutputMode {
				panic("Illegal output mode!")
			}
			currentPosition := d.position.Move(moveCommand)
			if tv, ok := d.grid[currentPosition]; ok && tv != tile {
				panic(fmt.Sprintf("%d != %d at %v", tv, tile, currentPosition))
			}
			d.grid[currentPosition] = tile
			if tile != wall {
				d.position = currentPosition
			}
		}
		for _, adjacent := range d.position.Adjacent() {
			if _, ok := d.grid[adjacent]; ok {
				continue
			}
			if _, ok := toVisitM[adjacent]; ok {
				continue
			}

			toVisit = append(toVisit, adjacent)
			toVisitM[adjacent] = true
		}
	}
}

type possibleMove struct {
	cord         Cord
	moveSequence []move
}

func (d *Droid) moveSequence(destination Cord) []move {
	position := d.position
	toVisit := []possibleMove{}
	toVisitM := map[Cord]bool{}
	for m, c := range position.Adjacent() {
		if c == destination {
			return []move{m}
		}
		if t, ok := d.grid[c]; !ok || t == wall {
			continue
		}
		toVisit = append(toVisit, possibleMove{c, []move{m}})
		toVisitM[c] = true
	}

	shortestPaths := map[Cord][]move{position: {}}

	for len(toVisit) > 0 {
		current := toVisit[0]
		delete(toVisitM, current.cord)
		toVisit = toVisit[1:]

		shortestPaths[current.cord] = current.moveSequence
		if current.cord == destination {
			return current.moveSequence
		}
		for m, nextCord := range current.cord.Adjacent() {
			if nextCord == destination {
				return append(current.moveSequence, m)
			}
			if _, ok := shortestPaths[nextCord]; ok {
				continue
			}
			if t, ok := d.grid[nextCord]; !ok || t == wall {
				continue
			}
			if _, ok := toVisitM[nextCord]; ok {
				continue
			}

			// here was 4 hour golang bug ;)
			// append(current.moveSequence, m)
			// modifies underliying array

			ms := make([]move, len(current.moveSequence))
			copy(ms, current.moveSequence)

			nextPossibleMove := possibleMove{nextCord, append(ms, m)}
			toVisit = append(toVisit, nextPossibleMove)
			toVisitM[nextCord] = true
		}
	}
	panic("not found!")
}

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)
	cpu := intcode.NewIntcode(intcode.NewMemory(lines[0]))
	d := NewDroid(cpu)
	d.Discover()
	d.position = Cord{0, 0}
	var oxygenCord Cord
	for c, t := range d.grid {
		if t == oxygenSystem {
			oxygenCord = c
			break
		}
	}
	fmt.Println(d.grid)
	fmt.Println("Exercise 1:", len(d.moveSequence(oxygenCord)))

	max := 0
	for c, t := range d.grid {
		if t == moved && c != oxygenCord {
			d.position = oxygenCord
			x := len(d.moveSequence(c))
			if x > max {
				max = x
			}
		}
	}

	fmt.Println("Exercise 2:", max)
}
