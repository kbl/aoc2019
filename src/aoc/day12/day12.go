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
	fmt.Printf("Exercise 2: %d\n", s2.Exercise2(s2.String()))
}

type Moon struct {
	x, y, z    int
	vx, vy, vz int
}

func (m *Moon) String() string {
	return fmt.Sprintf("pos=<x=%d, y=%d, z=%d>, vel=<x=%d, y=%d, z=%d>", m.x, m.y, m.z, m.vx, m.vy, m.vz)
}

func (s *Space) Exercise2(s2 string) int {
	initialEnergy := s.Energy()
	step := 0
	for {
		s.Simulate(1)
		if s.moons[0].Energy() == 0 {
			if s.Energy() == initialEnergy {
				fmt.Println(step)
				if s.String() == s2 {
					fmt.Println("DONE")
					fmt.Println(step + 1)
					return step
				}
			}
		}
		if step%500000000 == 0 {
			fmt.Println("mod", step)
		}
		step++
	}
	// 1703057968
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
	moons []*Moon
}

func NewSpace(s string) *Space {
	var m []*Moon
	for _, l := range strings.Split(s, "\n") {
		m = append(m, NewMoon(l))
	}
	return &Space{m}
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
		for _, m1 := range s.moons {
			for _, m2 := range s.moons {
				m1.vx += vel(m1.x, m2.x)
				m1.vy += vel(m1.y, m2.y)
				m1.vz += vel(m1.z, m2.z)
			}
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
