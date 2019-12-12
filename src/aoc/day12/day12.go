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
	allLines := strings.Join(lines, "\n")

	fmt.Printf("Exercise 1: %d\n", NewSpace(allLines).EnergyAfter(1000))
	fmt.Printf("Exercise 2: %d\n", NewSpace(allLines).PeriodLength())
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

func (s *Space) PeriodLength() int {
	xperiod := map[four]bool{}
	yperiod := map[four]bool{}
	zperiod := map[four]bool{}

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

func lcm(a, b, c int) int {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	lcmValue := 1
	for _, p := range primes(max) {
		canBeDivided := true
		for canBeDivided {
			canBeDivided = false
			if a%p == 0 {
				a /= p
				canBeDivided = true
			}
			if b%p == 0 {
				b /= p
				canBeDivided = true
			}
			if c%p == 0 {
				c /= p
				canBeDivided = true
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

func (s *Space) EnergyAfter(rounds int) int {
	s.Simulate(rounds)
	return s.Energy()
}
