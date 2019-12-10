package main

import (
	"aoc"
	"fmt"
	"sort"
	"strings"
)

type Cord struct {
	x, y float64
}

type Cords []Cord

func (c Cords) Len() int {
	return len(c)
}
func (c Cords) Less(i, j int) bool {
	if c[i].x == c[j].x {
		return c[i].y < c[j].y
	}
	return c[i].x < c[j].x
}
func (c Cords) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

type Func struct {
	a, b float64
}

func NewFunc(p1, p2 Cord) Func {
	// a*x1 - y1 + b = 0
	// a*x2 - y2 + b = 0
	// b = y1 - a*x1
	// a*x2 - y2 + y1 - a*x1 = 0
	// a(x2 - x1) = y2 - y1
	// a = (y2 - y1) / (x2 - x1)
	// b = y2 - a*x2
	a := (p2.y - p1.y) / (p2.x - p1.x)
	b := p2.y - a*p2.x
	return Func{a, b}
}

type Space struct {
	m         map[Cord]bool
	functions map[Func]Cords
	vertical  map[float64]Cords
}

func NewSpace(s string) *Space {
	m := map[Cord]bool{}
	for y, l := range strings.Split(s, "\n") {
		for x, a := range strings.Trim(l, " ") {
			if a == '#' {
				m[Cord{float64(x), float64(y)}] = true
			}
		}
	}
	functions, vertical := group(m)
	return &Space{m, functions, vertical}
}

func group(m map[Cord]bool) (map[Func]Cords, map[float64]Cords) {
	tempF := map[Func]map[Cord]bool{}
	tempV := map[float64]map[Cord]bool{}
	for p1 := range m {
		for p2 := range m {
			if p1 == p2 {
				continue
			}
			if p1.x == p2.x {
				if cords, ok := tempV[p1.x]; ok {
					cords[p1] = true
					cords[p2] = true
				} else {
					tempV[p1.x] = map[Cord]bool{p1: true, p2: true}
				}
				continue
			}

			f := NewFunc(p1, p2)

			if cords, ok := tempF[f]; ok {
				cords[p1] = true
				cords[p2] = true
			} else {
				tempF[f] = map[Cord]bool{p1: true, p2: true}
			}
		}
	}

	functions := map[Func]Cords{}
	vertical := map[float64]Cords{}

	for f, cords := range tempF {
		for c := range cords {
			functions[f] = append(functions[f], c)
		}
		sort.Sort(functions[f])
	}

	for x, cords := range tempV {
		for c := range cords {
			vertical[x] = append(vertical[x], c)
		}
		sort.Sort(vertical[x])
	}

	return functions, vertical
}

func (s *Space) Visible(c Cord) int {
	visible := map[Cord]bool{}
	for _, cords := range s.functions {
		for _, v := range pickVisible(c, cords) {
			visible[v] = true
		}
	}

	for _, cords := range s.vertical {
		for _, v := range pickVisible(c, cords) {
			visible[v] = true
		}
	}

	return len(visible)
}

func pickVisible(c Cord, cords Cords) Cords {
	visible := Cords{}
	cIndex := -1

	for i, oc := range cords {
		if oc == c {
			cIndex = i
			break
		}
	}

	if cIndex == -1 {
		return visible
	}

	if cIndex > 0 {
		visible = append(visible, cords[cIndex-1])
	}
	if cIndex < len(cords)-1 {
		visible = append(visible, cords[cIndex+1])
	}
	return visible
}

func main() {
	inputFilePath := aoc.InputArg()
	Main(inputFilePath)
}

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)
	linesStr := strings.Join(lines, "\n")
	s := NewSpace(linesStr)
	for c := range s.m {
		fmt.Printf("%v: %d\n", c, s.Visible(c))
	}
}
