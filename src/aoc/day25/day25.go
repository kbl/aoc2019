package main

import (
	"aoc"
	"aoc/intcode"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	inputFilePath := aoc.InputArg()
	Main(inputFilePath)
}

type cord struct {
	x, y int
}

var moves = map[string][]int{
	"north": {0, 1},
	"south": {0, -1},
	"east":  {1, 0},
	"west":  {-1, 0},
}

func (c cord) move(direction string) cord {
	return cord{
		c.x + moves[direction][0],
		c.y + moves[direction][1],
	}
}

var avoid_items = map[string]bool{"escape pod": true, "molten lava": true, "giant electromagnet": true, "infinite loop": true, "photons": true}

type room struct {
	name, description string
	doors             []string
	items             []string
}

func NewRoom(description string) *room {
	lines := strings.Split(strings.Trim(description, "\n"), "\n")
	r := room{}

	re, _ := regexp.Compile("== (.+) ==")
	r.name = string(re.FindSubmatch([]byte(lines[0]))[1])
	r.description = lines[1]
	lines = lines[3:]

	if strings.HasPrefix(lines[0], "Doors") {
		i := 1
		for strings.HasPrefix(lines[i], "- ") {
			r.doors = append(r.doors, strings.Split(lines[i], " ")[1])
			i++
		}
		lines = lines[i+1:]
	}

	if strings.HasPrefix(lines[0], "Items") {
		i := 1
		for strings.HasPrefix(lines[i], "- ") {
			r.items = append(r.items, strings.Split(lines[i], "- ")[1])
			i++
		}
		lines = lines[i+1:]
	}

	return &r
}

type gameSolver struct {
	mazeTraversed          bool
	rooms                  map[cord]*room
	minx, maxx, miny, maxy int
	position               cord
	inventory              map[string]bool
	toVisitRooms           []cord
}

func newGameSolver() *gameSolver {
	return &gameSolver{
		rooms:     map[cord]*room{},
		inventory: map[string]bool{},
	}
}

func (gs *gameSolver) visited(c cord) bool {
	_, ok := gs.rooms[c]
	return ok
}

func (gs *gameSolver) move(direction string) {
	for _, possibleDirection := range gs.rooms[gs.position].doors {
		if direction == possibleDirection {
			gs.position = gs.position.move(direction)
			break
		}
	}
}

func (gs *gameSolver) addRoom(c cord, r *room) {
	if c.x > gs.maxx {
		gs.maxx = c.x
	}
	if c.y > gs.maxy {
		gs.maxy = c.y
	}
	if c.x < gs.minx {
		gs.minx = c.x
	}
	if c.y < gs.miny {
		gs.miny = c.y
	}
	gs.rooms[c] = r

	newToVisitRooms := []cord{}
	for _, tvr := range gs.toVisitRooms {
		if c == tvr {
			continue
		}
		newToVisitRooms = append(newToVisitRooms, tvr)
	}
	gs.toVisitRooms = newToVisitRooms

	for _, direction := range r.doors {
		newPosition := gs.position.move(direction)
		if _, visited := gs.rooms[newPosition]; !visited {
			gs.toVisitRooms = append(gs.toVisitRooms, newPosition)
		}
	}
}

func (gs *gameSolver) findPath(start, end cord) []string {
	if !gs.visited(start) {
		panic("Path doesn't exist!")
	}
	if start == end {
		return nil
	}

	paths := map[cord][]cord{start: {}}
	directions := map[cord][]string{start: {}}

	toVisit := []cord{start}

	for len(toVisit) > 0 {
		current := toVisit[0]
		toVisit = toVisit[1:]

		for _, move := range gs.rooms[current].doors {
			next := current.move(move)
			if next == end {
				return append(directions[current], move)
			}
			if _, known := gs.rooms[next]; !known {
				continue
			}
			if _, alreadyTracked := paths[next]; alreadyTracked {
				continue
			}

			toVisit = append(toVisit, next)

			paths[next] = append([]cord{}, paths[current]...)
			paths[next] = append(paths[next], next)

			directions[next] = append([]string{}, directions[current]...)
			directions[next] = append(directions[next], move)
		}
	}

	return directions[end]
}

func (gs *gameSolver) String() string {
	lines := []string{}
	for y := gs.maxy + 1; y >= gs.miny-1; y-- {
		line := []rune{}
		for x := gs.minx - 1; x <= gs.maxx+1; x++ {
			character := ' '
			c := cord{x, y}
			if _, ok := gs.rooms[c]; ok {
				if x == 0 && y == 0 {
					character = '®'
				} else if x == gs.position.x && y == gs.position.y {
					character = 'X'
				} else {
					character = '.'
				}
			}
			for _, tvr := range gs.toVisitRooms {
				if tvr == c {
					character = '?'
				}
			}
			line = append(line, character)
		}
		lines = append(lines, string(line))
	}
	return strings.Join(lines, "\n")
}

func (gs *gameSolver) takeItems(cpu *intcode.Intcode) {
	for _, i := range gs.rooms[gs.position].items {
		if avoid_items[i] {
			continue
		}
		fmt.Println(">> taking", i)
		gs.inventory[i] = true
		cpu.AddAsciiInput(fmt.Sprintf("take %s", i))
	}
}

func (gs *gameSolver) pickNextRoom(cpu *intcode.Intcode) {
	if len(gs.toVisitRooms) == 0 {
		gs.mazeTraversed = true
		return
	}
	if gs.position == gs.toVisitRooms[0] {
		gs.toVisitRooms = gs.toVisitRooms[1:]
	}

	path := gs.findPath(gs.position, gs.toVisitRooms[0])
	fmt.Println(">> going", path[0])
	fmt.Println(gs.position)
	cpu.AddAsciiInput(path[0])
	gs.move(path[0])
}

type outputType int

const (
	outputRoom        = 0
	outputTakeItem    = 1
	outputMissingItem = 2
	outputDropItem    = 3
	outputInventory   = 4
)

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)
	cpu := intcode.NewIntcode(intcode.NewMemory(lines[0]))
	gs := newGameSolver()

	output := []rune{}
	for !gs.mazeTraversed {
		value, mode := cpu.Output()
		if mode == intcode.HaltMode {
			fmt.Println(value)
			panic("HALT MODE")
		}
		output = append(output, rune(value))
		if strings.HasSuffix(string(output), "Command?\n") {
			// fmt.Println(string(output))
			switch decide(string(output)) {
			case outputRoom:
				room := NewRoom(string(output))
				gs.addRoom(gs.position, room)
				gs.takeItems(cpu)
				gs.pickNextRoom(cpu)
			case outputTakeItem:
				break
			default:
				panic("unknown output")
			}
			fmt.Println(gs)
			output = []rune{}
		}
	}

	output = []rune{}

	items := []string{}
	for i := range gs.inventory {
		items = append(items, i)
	}
	sort.Strings(items)
	combs := combinations(items)
	for _, c := range combs {
		fmt.Println(c)
	}

	for _, i := range items {
		cpu.AddAsciiInput(fmt.Sprintf("drop %s", i))
	}

	dropped := 0
	taken := 0
	for combinationId := 0; combinationId < len(combs); {
		value, mode := cpu.Output()
		if mode == intcode.HaltMode {
			fmt.Println(value)
			fmt.Println(string(output))
			return
		}
		output = append(output, rune(value))
		if strings.HasSuffix(string(output), "Command?\n") {
			// fmt.Println(string(output))
			switch decide(string(output)) {
			case outputMissingItem:
				dropped++
			case outputDropItem:
				dropped++
			case outputTakeItem:
				taken++
			case outputRoom:
				fmt.Println(string(output))
				for _, i := range items {
					cpu.AddAsciiInput(fmt.Sprintf("drop %s", i))
				}
				combinationId++
			case outputInventory:
				fmt.Println(string(output))
			default:
				fmt.Println(string(output))
				panic("Unknown output!")
			}

			if dropped == len(items) {
				fmt.Println(combinationId)
				for _, i := range combs[combinationId] {
					cpu.AddAsciiInput(fmt.Sprintf("take %s", i))
				}
				dropped = 0
			}

			if taken == len(combs[combinationId]) {
				cpu.AddAsciiInput("inv")
				cpu.AddAsciiInput("west")
				taken = 0
			}
			output = []rune{}
		}
	}
}

func combinations(items []string) [][]string {
	result := [][]string{}
	for i := 1; i < 1<<len(items); i++ {
		combination := []string{}
		for j := 0; j < len(items); j++ {
			if (1 << j & i) > 0 {
				combination = append(combination, items[j])
			}
		}
		result = append(result, combination)
	}
	return result
}

func (gs *gameSolver) manual(cpu *intcode.Intcode) {
	reader := bufio.NewReader(os.Stdin)
	command, _ := reader.ReadString('\n')
	command = strings.Trim(command, "\n")
	if _, ok := moves[command]; ok {
		gs.move(command)
	}
	cpu.AddAsciiInput(command)
}

func decide(output string) outputType {
	if strings.Contains(output, "You take ") {
		return outputTakeItem
	}
	if strings.Contains(output, "== ") {
		return outputRoom
	}
	if strings.Contains(output, "You don't have that item.") {
		return outputMissingItem
	}
	if strings.Contains(output, "You drop the") {
		return outputDropItem
	}
	if strings.Contains(output, "Items in your inventory") {
		return outputInventory
	}
	return -1
}
