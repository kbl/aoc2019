package main

import (
	"aoc"
	"aoc/intcode"
	"bufio"
	"fmt"
	"os"
	"regexp"
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

var avoid_items = []string{"escape pod", "molten lava", "giant electromagnet", "infinite loop", "photons"}

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
			r.items = append(r.items, strings.Split(lines[i], " ")[1])
			i++
		}
		lines = lines[i+1:]
	}

	return &r
}

type maze struct {
	rooms                  map[cord]*room
	minx, maxx, miny, maxy int
	position               cord
	toVisit                map[cord]bool
}

func NewMaze() *maze {
	return &maze{
		rooms:   make(map[cord]*room),
		toVisit: make(map[cord]bool),
	}
}

func (m *maze) visited(c cord) bool {
	_, ok := m.rooms[c]
	return ok
}

func (m *maze) move(direction string) {
	for _, possibleDirection := range m.rooms[m.position].doors {
		if direction == possibleDirection {
			m.position = m.position.move(direction)
			break
		}
	}
}

func (m *maze) addRoom(c cord, r *room) {
	if c.x > m.maxx {
		m.maxx = c.x
	}
	if c.y > m.maxy {
		m.maxy = c.y
	}
	if c.x < m.minx {
		m.minx = c.x
	}
	if c.y < m.miny {
		m.miny = c.y
	}
	m.rooms[c] = r
	m.toVisit[c] = false
	for _, direction := range r.doors {
		newPosition := m.position.move(direction)
		if _, visited := m.rooms[newPosition]; !visited {
			m.toVisit[newPosition] = true
		}
	}
}

func (m *maze) findPath(start, end cord) []string {
	if !m.visited(start) {
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

		for move := range moves {
			next := current.move(move)
			if next == end {
				return append(directions[current], move)
			}
			if _, known := m.rooms[next]; !known {
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

func (m *maze) String() string {
	lines := []string{}
	for y := m.maxy + 1; y >= m.miny-1; y-- {
		line := []rune{}
		for x := m.minx - 1; x <= m.maxx+1; x++ {
			if _, ok := m.rooms[cord{x, y}]; ok {
				if x == m.position.x && y == m.position.y {
					line = append(line, 'X')
				} else {
					line = append(line, '.')
				}
			} else if m.toVisit[cord{x, y}] {
				line = append(line, '?')
			} else {
				line = append(line, ' ')
			}
		}
		lines = append(lines, string(line))
	}
	return strings.Join(lines, "\n")
}

type outputType int

const (
	outputRoom = 0
	outputTake = 1
)

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)
	cpu := intcode.NewIntcode(intcode.NewMemory(lines[0]))

	grid := NewMaze()

	output := []rune{}
	reader := bufio.NewReader(os.Stdin)
	for {
		value, mode := cpu.Output()
		if mode == intcode.HaltMode {
			panic("HALT MODE")
		}
		output = append(output, rune(value))
		if strings.HasSuffix(string(output), "Command?\n") {
			fmt.Println(string(output))
			switch decide(string(output)) {
			case outputRoom:
				room := NewRoom(string(output))
				grid.addRoom(grid.position, room)
			}
			fmt.Println(grid)
			command, _ := reader.ReadString('\n')
			command = strings.Trim(command, "\n")
			if _, ok := moves[command]; ok {
				grid.move(command)
			}
			cpu.AddAsciiInput(command)
			output = []rune{}
		}
	}
}

func decide(output string) outputType {
	if strings.Contains(output, "You take ") {
		return outputTake
	}
	if strings.Contains(output, "== ") {
		return outputRoom
	}
	return -1
}
