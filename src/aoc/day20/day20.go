package main

import (
	"aoc"
	"fmt"
	"strings"
)

const frame = 2
const wall = '#'
const path = '.'

func main() {
	inputFilePath := aoc.InputArg()
	Main(inputFilePath)
}

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)
	fmt.Printf("Read %d lines!\n", len(lines))
}

type Vertex string
type Edge [2]Vertex

type cord struct {
	x, y int
}

func (c cord) adjacent() []cord {
	a := []cord{
		cord{c.x + 1, c.y},
		cord{c.x, c.y + 1},
	}
	if c.x > 0 {
		a = append(a, cord{c.x - 1, c.y})
	}
	if c.y > 0 {
		a = append(a, cord{c.x, c.y - 1})
	}
	return a
}

type Graph struct {
	Edges     map[Edge]int
	Vertices  map[Vertex]bool
	entrances map[Vertex][]cord
}

type maze struct {
	m          map[cord]rune
	lenX, lenY int
}

func newMaze(mazeStr string) *maze {
	m := maze{m: map[cord]rune{}}

	for y, rawRow := range strings.Split(mazeStr, "\n") {
		for x, tile := range rawRow {
			m.m[cord{x, y}] = tile
			m.lenX = x + 1
		}
		m.lenY = y + 1
	}

	return &m
}

func (m *maze) findDonutWidth() int {
	c := cord{m.lenX / 2, m.lenY / 2}
	for {
		switch m.m[c] {
		case '.':
		case '#':
			return c.x + 1 - frame
		}
		c = cord{c.x - 1, c.y}
	}
}

func (m *maze) findEntrances() map[Vertex][]cord {
	dw := m.findDonutWidth()
	entrances := map[Vertex][]cord{}

	// entrances on horizontal top outer edge
	y := frame
	for x := frame; x < m.lenX-frame; x++ {
		if m.m[cord{x, y}] == path {
			v := Vertex([]rune{m.m[cord{x, y - 2}], m.m[cord{x, y - 1}]})
			entrances[v] = append(entrances[v], cord{x, y})
		}
	}

	// entrances on horizontal bottom outer edge
	y = m.lenY - frame - 1
	for x := frame; x < m.lenX-frame; x++ {
		if m.m[cord{x, y}] == path {
			v := Vertex([]rune{m.m[cord{x, y + 1}], m.m[cord{x, y + 2}]})
			entrances[v] = append(entrances[v], cord{x, y})
		}
	}

	// entrances on vertical left outer edge
	x := frame
	for y := frame; y < m.lenY-frame; y++ {
		if m.m[cord{x, y}] == path {
			v := Vertex([]rune{m.m[cord{x - 2, y}], m.m[cord{x - 1, y}]})
			entrances[v] = append(entrances[v], cord{x, y})
		}
	}

	// entrances on vertical right outer edge
	x = m.lenX - frame - 1
	for y := frame; y < m.lenY-frame; y++ {
		if m.m[cord{x, y}] == path {
			v := Vertex([]rune{m.m[cord{x + 1, y}], m.m[cord{x + 2, y}]})
			entrances[v] = append(entrances[v], cord{x, y})
		}
	}

	// entrances on horizontal top inner edge
	y = frame + dw - 1
	for x := frame + dw; x < m.lenX-frame-dw; x++ {
		if m.m[cord{x, y}] == path {
			v := Vertex([]rune{m.m[cord{x, y + 1}], m.m[cord{x, y + 2}]})
			entrances[v] = append(entrances[v], cord{x, y})
		}
	}

	// entrances on horizontal bottom inner edge
	y = m.lenY - frame - dw
	for x := frame + dw; x < m.lenX-frame-dw; x++ {
		if m.m[cord{x, y}] == path {
			v := Vertex([]rune{m.m[cord{x, y - 2}], m.m[cord{x, y - 1}]})
			entrances[v] = append(entrances[v], cord{x, y})
		}
	}

	// entrances on vertical left inner edge
	x = frame + dw - 1
	for y := frame + dw; y < m.lenY-frame-dw; y++ {
		if m.m[cord{x, y}] == path {
			v := Vertex([]rune{m.m[cord{x + 1, y}], m.m[cord{x + 2, y}]})
			entrances[v] = append(entrances[v], cord{x, y})
		}
	}

	// entrances on vertical right inner edge
	x = m.lenX - frame - dw
	for y := frame + dw; y < m.lenY-frame-dw; y++ {
		if m.m[cord{x, y}] == path {
			v := Vertex([]rune{m.m[cord{x - 2, y}], m.m[cord{x - 1, y}]})
			entrances[v] = append(entrances[v], cord{x, y})
		}
	}

	return entrances
}

func NewGraph(maze string) *Graph {
	m := newMaze(maze)

	g := Graph{
		Edges:     map[Edge]int{},
		Vertices:  map[Vertex]bool{},
		entrances: m.findEntrances(),
	}

	ctv := map[cord]Vertex{}

	for v, ends := range m.findEntrances() {
		g.Vertices[v] = true
		for _, e := range ends {
			ctv[e] = v
		}
	}

	for start, vertex := range ctv {
		for otherVertex, length := range findEdges(m.m, start, vertex, ctv) {
			edge := Edge{vertex, otherVertex}
			if otherVertex > vertex {
				edge = Edge{otherVertex, vertex}
			}
			if e, ok := g.Edges[edge]; ok && e != length {
				panic("Those lenghts should be equal!")
			}
			g.Edges[edge] = length
		}
	}

	return &g
}

func findEdges(g map[cord]rune, start cord, v1 Vertex, ctv map[cord]Vertex) map[Vertex]int {
	distances := map[cord]int{start: 1}
	toVisit := []cord{start}
	for len(toVisit) > 0 {
		current := toVisit[0]
		toVisit = toVisit[1:]

		fmt.Println(distances, current)
	}
	return map[Vertex]int{}
}
