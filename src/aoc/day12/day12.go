package main

import (
	"aoc"
	"fmt"
	"log"
	"strings"
)

const debug = false

func main() {
	inputFilePath := aoc.InputArg()
	Main(inputFilePath)
}

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)
	allLines := strings.Join(lines, "\n")

	fmt.Printf("Exercise 1: %d\n", NewSpace(allLines).EnergyAfter(1000))
	fmt.Printf("Exercise 2: %d\n", NewSpace(allLines).PeriodLength())
}

type Moon struct {
	position [3]int
	velocity [3]int
}

func (m *Moon) String() string {
	return fmt.Sprintf("pos=<x=%d, y=%d, z=%d>, vel=<x=%d, y=%d, z=%d>", m.position[0], m.position[1], m.position[2], m.velocity[0], m.velocity[1], m.velocity[2])
}

type eight struct {
	a, av, b, bv, c, cv, d, dv int
}

func (s *Space) PeriodLength() int {
	xperiod := map[eight]bool{}
	yperiod := map[eight]bool{}
	zperiod := map[eight]bool{}

	fmaker := func(dimension int) func() eight {
		return func() eight {
			return eight{
				s.moons[0].position[dimension],
				s.moons[0].velocity[dimension],
				s.moons[1].position[dimension],
				s.moons[1].velocity[dimension],
				s.moons[2].position[dimension],
				s.moons[2].velocity[dimension],
				s.moons[3].position[dimension],
				s.moons[3].velocity[dimension],
			}
		}
	}

	fx := fmaker(0)
	fy := fmaker(1)
	fz := fmaker(2)

	xperiod[fx()] = true
	yperiod[fy()] = true
	zperiod[fz()] = true

	seen_x := false
	seen_y := false
	seen_z := false

	for !(seen_x && seen_y && seen_z) {
		s.Simulate(1)
		x := fx()
		y := fy()
		z := fz()

		if _, ok := xperiod[x]; ok {
			seen_x = true
		} else {
			xperiod[x] = true
		}
		if _, ok := yperiod[y]; ok {
			seen_y = true
		} else {
			yperiod[y] = true
		}
		if _, ok := zperiod[z]; ok {
			seen_z = true
		} else {
			zperiod[z] = true
		}
	}

	return lcm(len(xperiod), len(yperiod), len(zperiod))
}

func lcm(values ...int) int {
	max := values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	lcmValue := 1
	for _, p := range primes(max / 2) {
		canBeDivided := true
		for canBeDivided {
			canBeDivided = false
			for i := range values {
				if values[i]%p == 0 {
					values[i] /= p
					canBeDivided = true
				}
			}
			if canBeDivided {
				lcmValue *= p
			}
		}
	}
	return lcmValue
}

func primes(max int) []int {
	p := []bool{}
	for i := 0; i <= max; i++ {
		p = append(p, true)
	}
	p[0] = false
	p[1] = false
	for i := 2; i <= max; i++ {
		if p[i] {
			for j := i * 2; j <= max; j += i {
				p[j] = false
			}
		}
	}
	primes := []int{}
	for i, v := range p {
		if v {
			primes = append(primes, i)
		}
	}
	return primes
}

func NewMoon(l string) *Moon {
	moon := Moon{}
	_, err := fmt.Sscanf(l, "<x=%d, y=%d, z=%d>", &moon.position[0], &moon.position[1], &moon.position[2])
	if err != nil {
		log.Fatal(err)
	}
	return &moon
}

func (m *Moon) Move() {
	for dimension := 0; dimension < 3; dimension++ {
		m.position[dimension] += m.velocity[dimension]
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (m *Moon) Energy() int {
	return (abs(m.position[0]) + abs(m.position[1]) + abs(m.position[2])) * (abs(m.velocity[0]) + abs(m.velocity[1]) + abs(m.velocity[2]))
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

			v := vel(m1.position[0], m2.position[0])
			m1.velocity[0] += v
			m2.velocity[0] -= v

			v = vel(m1.position[1], m2.position[1])
			m1.velocity[1] += v
			m2.velocity[1] -= v

			v = vel(m1.position[2], m2.position[2])
			m1.velocity[2] += v
			m2.velocity[2] -= v
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

func (s *Space) EnergyAfter(rounds int) int {
	s.Simulate(rounds)
	return s.Energy()
}
