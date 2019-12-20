package main

import (
	"aoc"
	"fmt"
	"sort"
	"strings"
)

const frame = 2
const wall = '#'
const path = '.'

const outer = 0
const inner = 1

func main() {
	inputFilePath := aoc.InputArg()
	Main(inputFilePath)
}

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)
	g := NewGraph(strings.Join(lines, "\n"))
	fmt.Println("Exercise 1:", g.ShortestPath(Vertex{"AA", outer}, Vertex{"ZZ", outer}))
	fmt.Println("Exercise 2:", g.RecursiveShortestPath(Vertex{"AA", outer}, Vertex{"ZZ", outer}))
}

type side int

type Vertex struct {
	v string
	s side
}
type Edge struct {
	v1, v2 Vertex
}

type cord struct {
	x, y int
}

func (c cord) adjacent() []cord {
	return []cord{
		{c.x + 1, c.y},
		{c.x, c.y + 1},
		{c.x - 1, c.y},
		{c.x, c.y - 1},
	}
}

type Graph struct {
	Edges      map[Edge]int
	dummyEdges map[string]map[Vertex]int
	Vertices   map[string]bool
	entrances  map[Vertex]cord
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

func (m *maze) findEntrances() map[Vertex]cord {
	dw := m.findDonutWidth()
	entrances := map[Vertex]cord{}

	// entrances on horizontal top outer edge
	y := frame
	for x := frame; x < m.lenX-frame; x++ {
		if m.m[cord{x, y}] == path {
			v := Vertex{string([]rune{m.m[cord{x, y - 2}], m.m[cord{x, y - 1}]}), outer}
			entrances[v] = cord{x, y}
		}
	}

	// entrances on horizontal bottom outer edge
	y = m.lenY - frame - 1
	for x := frame; x < m.lenX-frame; x++ {
		if m.m[cord{x, y}] == path {
			v := Vertex{string([]rune{m.m[cord{x, y + 1}], m.m[cord{x, y + 2}]}), outer}
			entrances[v] = cord{x, y}
		}
	}

	// entrances on vertical left outer edge
	x := frame
	for y := frame; y < m.lenY-frame; y++ {
		if m.m[cord{x, y}] == path {
			v := Vertex{string([]rune{m.m[cord{x - 2, y}], m.m[cord{x - 1, y}]}), outer}
			entrances[v] = cord{x, y}
		}
	}

	// entrances on vertical right outer edge
	x = m.lenX - frame - 1
	for y := frame; y < m.lenY-frame; y++ {
		if m.m[cord{x, y}] == path {
			v := Vertex{string([]rune{m.m[cord{x + 1, y}], m.m[cord{x + 2, y}]}), outer}
			entrances[v] = cord{x, y}
		}
	}

	// entrances on horizontal top inner edge
	y = frame + dw - 1
	for x := frame + dw; x < m.lenX-frame-dw; x++ {
		if m.m[cord{x, y}] == path {
			v := Vertex{string([]rune{m.m[cord{x, y + 1}], m.m[cord{x, y + 2}]}), inner}
			entrances[v] = cord{x, y}
		}
	}

	// entrances on horizontal bottom inner edge
	y = m.lenY - frame - dw
	for x := frame + dw; x < m.lenX-frame-dw; x++ {
		if m.m[cord{x, y}] == path {
			v := Vertex{string([]rune{m.m[cord{x, y - 2}], m.m[cord{x, y - 1}]}), inner}
			entrances[v] = cord{x, y}
		}
	}

	// entrances on vertical left inner edge
	x = frame + dw - 1
	for y := frame + dw; y < m.lenY-frame-dw; y++ {
		if m.m[cord{x, y}] == path {
			v := Vertex{string([]rune{m.m[cord{x + 1, y}], m.m[cord{x + 2, y}]}), inner}
			entrances[v] = cord{x, y}
		}
	}

	// entrances on vertical right inner edge
	x = m.lenX - frame - dw
	for y := frame + dw; y < m.lenY-frame-dw; y++ {
		if m.m[cord{x, y}] == path {
			v := Vertex{string([]rune{m.m[cord{x - 2, y}], m.m[cord{x - 1, y}]}), inner}
			entrances[v] = cord{x, y}
		}
	}

	return entrances
}

func NewGraph(maze string) *Graph {
	m := newMaze(maze)

	g := Graph{
		Edges:      map[Edge]int{},
		dummyEdges: map[string]map[Vertex]int{},
		Vertices:   map[string]bool{},
		entrances:  m.findEntrances(),
	}

	ctv := map[cord]Vertex{}

	for v, end := range m.findEntrances() {
		g.Vertices[v.v] = true
		ctv[end] = v
	}

	for start, vertex := range ctv {
		for otherVertex, length := range findEdges(m.m, start, vertex, ctv) {
			edge := Edge{vertex, otherVertex}
			if otherVertex.v > vertex.v {
				edge = Edge{otherVertex, vertex}
			}
			if e, ok := g.Edges[edge]; ok && e != length {
				panic("Those lenghts should be equal!")
			}
			g.Edges[edge] = length
			if m, ok := g.dummyEdges[vertex.v]; ok {
				m[otherVertex] = length
			} else {
				g.dummyEdges[vertex.v] = map[Vertex]int{otherVertex: length}
			}
		}
	}

	return &g
}

func findEdges(g map[cord]rune, start cord, v1 Vertex, ctv map[cord]Vertex) map[Vertex]int {
	edges := map[Vertex]int{}
	distances := map[cord]int{start: 1}
	toVisit := []cord{start}

	for len(toVisit) > 0 {
		current := toVisit[0]
		toVisit = toVisit[1:]

		for _, a := range current.adjacent() {
			valid := g[a] == path
			for _, tv := range toVisit {
				if a == tv {
					valid = false
					break
				}
			}
			if _, ok := distances[a]; ok {
				valid = false
			}

			if valid {
				toVisit = append(toVisit, a)
				distances[a] = distances[current] + 1
			}
		}

		if v2, ok := ctv[current]; ok && v2 != v1 {
			edges[v2] = distances[current]
			continue
		}
	}
	return edges
}

type trail struct {
	end      Vertex
	distance int
}
type trails []trail

func (t trails) Len() int {
	return len(t)
}

func (t trails) Less(i, j int) bool {
	return t[i].distance < t[j].distance
}

func (t trails) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (g *Graph) ShortestPath(start, end Vertex) int {
	visited := map[Vertex]int{start: 0}
	queue := trails{}
	for v, l := range g.dummyEdges[start.v] {
		queue = append(queue, trail{v, l})
	}

	for len(queue) > 0 {
		sort.Sort(queue)
		current := queue[0]
		queue = queue[1:]

		if _, ok := visited[current.end]; ok {
			continue
		}
		visited[current.end] = current.distance

		for v, l := range g.dummyEdges[current.end.v] {
			queue = append(queue, trail{v, current.distance + l})
		}
	}

	return visited[end] - 1
}

func (g *Graph) RecursiveShortestPath(start, end Vertex) int {
	visited := map[Vertex]int{start: 0}
	queue := trails{}
	for v, l := range g.dummyEdges[start.v] {
		queue = append(queue, trail{v, l})
	}

	for len(queue) > 0 {
		sort.Sort(queue)
		current := queue[0]
		queue = queue[1:]

		if _, ok := visited[current.end]; ok {
			continue
		}
		visited[current.end] = current.distance

		for v, l := range g.dummyEdges[current.end.v] {
			queue = append(queue, trail{v, current.distance + l})
		}
	}

	return visited[end] - 1
}
