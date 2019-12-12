package main

import (
	"aoc"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

const debug = false

func main() {
	inputFilePath := aoc.InputArg()
	Main(inputFilePath)
}

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)
	s := NewSpace(strings.Join(lines, "\n"))
	s2 := NewSpace(strings.Join(lines, "\n"))
	s.Simulate(1000)
	fmt.Printf("Exercise 1: %d\n", s.Energy())
	fmt.Printf("Exercise 2: %d\n", s2.Exercise2())
}

type Moon struct {
	x, y, z    int
	vx, vy, vz int
}

func (m *Moon) String() string {
	return fmt.Sprintf("pos=<x=%d, y=%d, z=%d>, vel=<x=%d, y=%d, z=%d>", m.x, m.y, m.z, m.vx, m.vy, m.vz)
}

type four struct {
	a, b, c, d, va, vb, vc, vd int
}

func (s *Space) Exercise2() int {
	seenx := map[four]int{}
	seeny := map[four]int{}
	seenz := map[four]int{}
	seen_x := false
	seen_y := false
	seen_z := false

	fx := func() four {
		return four{
			s.moons[0].x,
			s.moons[1].x,
			s.moons[2].x,
			s.moons[3].x,
			s.moons[0].vx,
			s.moons[1].vx,
			s.moons[2].vx,
			s.moons[3].vx,
		}
	}
	fy := func() four {
		return four{
			s.moons[0].y,
			s.moons[1].y,
			s.moons[2].y,
			s.moons[3].y,
			s.moons[0].vy,
			s.moons[1].vy,
			s.moons[2].vy,
			s.moons[3].vy,
		}
	}
	fz := func() four {
		return four{
			s.moons[0].z,
			s.moons[1].z,
			s.moons[2].z,
			s.moons[3].z,
			s.moons[0].vz,
			s.moons[1].vz,
			s.moons[2].vz,
			s.moons[3].vz,
		}
	}

	seenx[fx()] = 0
	seeny[fy()] = 0
	seenz[fz()] = 0

	step := 0
	for {
		s.Simulate(1)
		step++
		x := fx()
		y := fy()
		z := fz()
		if _, ok := seenx[x]; ok {
			seen_x = true
		} else {
			seen_x = false
			seenx[x] = step
		}
		if _, ok := seeny[y]; ok {
			seen_y = true
		} else {
			seen_y = false
			seeny[y] = step
		}
		if _, ok := seenz[z]; ok {
			seen_z = true
		} else {
			seen_z = false
			seenz[z] = step
		}

		fmt.Println(seen_x, seen_y, seen_z, step)

		if seen_x && seen_y && seen_z {
			break
		}
	}

	fmt.Println(len(seenx))
	fmt.Println(len(seeny))
	fmt.Println(len(seenz))

	return 0
}

func NewMoon(l string) *Moon {
	t := strings.Split(l, "=")
	x, err := strconv.Atoi(strings.Split(t[1], ",")[0])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(strings.Split(t[2], ",")[0])
	if err != nil {
		log.Fatal(err)
	}
	z, err := strconv.Atoi(strings.Split(t[3], ">")[0])
	if err != nil {
		log.Fatal(err)
	}
	return &Moon{x: x, y: y, z: z}
}

func (m *Moon) Move() {
	m.x += m.vx
	m.y += m.vy
	m.z += m.vz
}

func (m *Moon) Energy() int {
	return int(
		math.Abs(float64(m.x))+
			math.Abs(float64(m.y))+
			math.Abs(float64(m.z))) * int(math.Abs(float64(m.vx))+
		math.Abs(float64(m.vy))+
		math.Abs(float64(m.vz)))
}

type Space struct {
	moons     []*Moon
	moonPairs [][]*Moon
}

func NewSpace(s string) *Space {
	var m []*Moon
	for _, l := range strings.Split(s, "\n") {
		m = append(m, NewMoon(l))
	}

	var mp [][]*Moon
	for i := range m {
		for j := i + 1; j < len(m); j++ {
			mp = append(mp, []*Moon{m[i], m[j]})
		}
	}

	return &Space{m, mp}
}

func vel(a, b int) int {
	if a == b {
		return 0
	}
	if a > b {
		return -1
	}
	return 1
}

func (s *Space) String() string {
	var moons []string
	for _, m := range s.moons {
		moons = append(moons, m.String())
	}
	return strings.Join(moons, "\n")
}

func (s *Space) Simulate(steps int) {
	if debug {
		fmt.Println("After 0 steps:")
		fmt.Println(s)
		fmt.Println()
	}
	for i := 1; i <= steps; i++ {
		for _, mp := range s.moonPairs {
			m1 := mp[0]
			m2 := mp[1]

			v := vel(m1.x, m2.x)
			m1.vx += v
			m2.vx -= v

			v = vel(m1.y, m2.y)
			m1.vy += v
			m2.vy -= v

			v = vel(m1.z, m2.z)
			m1.vz += v
			m2.vz -= v
		}

		for _, m := range s.moons {
			m.Move()
		}

		if debug {
			fmt.Printf("After %d steps:\n", i)
			fmt.Println(s)
			fmt.Println()
		}
	}
}

func (s *Space) Energy() int {
	te := 0
	for _, m := range s.moons {
		te += m.Energy()
	}
	return te
}
