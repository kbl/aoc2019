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

type Funcs []Func

func (f Funcs) Len() int {
	return len(f)
}
func (f Funcs) Less(i, j int) bool {
	if f[i].a == f[j].a {
		return f[i].b < f[j].b
	}
	return f[i].a < f[j].a
}
func (f Funcs) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
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

func group2(c Cord, m map[Cord]bool) (map[Func]Cords, map[float64]Cords) {
	tempF := map[Func]map[Cord]bool{}
	tempV := map[float64]map[Cord]bool{}
	for p1 := range m {
		if c == p1 {
			continue
		}
		if p1.x == c.x {
			if cords, ok := tempV[p1.x]; ok {
				cords[p1] = true
				cords[c] = true
			} else {
				tempV[p1.x] = map[Cord]bool{p1: true, c: true}
			}
			continue
		}

		f := NewFunc(p1, c)

		if cords, ok := tempF[f]; ok {
			cords[p1] = true
			cords[c] = true
		} else {
			tempF[f] = map[Cord]bool{p1: true, c: true}
		}
	}

	functions := map[Func]Cords{}
	vertical := map[float64]Cords{}

	for f, cords := range tempF {
		for cc := range cords {
			functions[f] = append(functions[f], cc)
		}
		sort.Sort(functions[f])
	}

	for x, cords := range tempV {
		for cc := range cords {
			vertical[x] = append(vertical[x], cc)
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

func scale(c Cord, m map[Cord]bool) map[Cord]bool {
	scaled := map[Cord]bool{}
	for oc := range m {
		scaled[Cord{oc.x - c.x, oc.y - c.y}] = true
	}
	return scaled
}

func (s *Space) Vaporize(c Cord, count int) Cord {
	m := s.m
	center := c
	functions, vertical := group2(center, m)
	funcMap := map[Func][]Cords{}
	funcs := Funcs{}
	for f, cords := range functions {
		before, after := split(center, cords)
		if len(before) > 0 || len(after) > 0 {
			funcMap[f] = []Cords{before, after}
			funcs = append(funcs, f)
		}
	}

	verts := []Cords{{}, {}}
	for _, cords := range vertical {
		before, after := split(center, cords)
		if len(before) > 0 || len(after) > 0 {
			verts = []Cords{before, after}
		}
	}

	rr := []Cords{}

	sort.Sort(funcs)

	for _, f := range funcs {
		fmt.Println(f, funcMap[f])
	}

	reverse(verts[0])
	rr = append(rr, verts[0])
	fmt.Println("v", "before", rr)
	fmt.Println(rr)

	for i := 0; i < len(funcs); i++ {
		f := funcs[i]
		cords := funcMap[f]
		rr = append(rr, cords[1])
		fmt.Println("f", f, "after", rr)
	}

	rr = append(rr, verts[1])
	fmt.Println("v", "after", rr)
	fmt.Println("v", "after", rr)
	fmt.Println("v", "after", rr)
	fmt.Println("v", "after", rr)

	for i := 0; i < len(funcs); i++ {
		f := funcs[i]
		cords := funcMap[f]
		reverse(cords[0])
		rr = append(rr, cords[0])
		fmt.Println("f", f, "before", rr)
	}

	fmt.Println(rr)

	alreadyVaporized := map[Cord]bool{}
	vaporized := 0
	i := 0
	for vaporized < count {
		for len(rr[i]) > 0 {
			toVaporize := rr[i][0]
			if _, ok := alreadyVaporized[toVaporize]; ok {
				rr[i] = rr[i][1:]
			} else {
				vaporized++
				alreadyVaporized[toVaporize] = true
				rr[i] = rr[i][1:]
				fmt.Println("Vaporizing", vaporized, toVaporize)
				break
			}
		}
		i = (i + 1) % len(rr)
	}

	// 401 too low
	// 219 too low

	return Cord{0, 0}
}

func reverse(cords Cords) {
	for i := 0; i < len(cords)/2; i++ {
		cords[i], cords[len(cords)-i-1] = cords[len(cords)-i-1], cords[i]
	}
}

func split(c Cord, cords Cords) (Cords, Cords) {
	cIndex := -1

	for i, oc := range cords {
		if oc == c {
			cIndex = i
			break
		}
	}

	if cIndex == -1 {
		return nil, nil
	}

	return cords[:cIndex], cords[cIndex+1:]
}

func pickVisible(c Cord, cords Cords) Cords {
	visible := Cords{}
	before, after := split(c, cords)
	if len(before) > 0 {
		visible = append(visible, before[len(before)-1])
	}
	if len(after) > 0 {
		visible = append(visible, after[0])
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

	s.Vaporize(Cord{11, 11}, 200)
}
