package main

import (
	"aoc"
	"aoc/collections"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	cut  = 0
	deal = 1
	inc  = 2
)

func parse(instructions []string) [][]int {
	ops := [][]int{}
	for _, l := range instructions {
		t := strings.Split(l, " ")
		if t[0] == "cut" {
			v, err := strconv.Atoi(t[1])
			if err != nil {
				log.Fatal(err)
			}
			ops = append(ops, []int{cut, v})
		} else if t[1] == "with" {
			v, err := strconv.Atoi(t[3])
			if err != nil {
				log.Fatal(err)
			}
			ops = append(ops, []int{inc, v})
		} else if t[1] == "into" {
			ops = append(ops, []int{deal, 0})
		} else {
			panic("unknown")
		}
	}
	return ops
}

func main() {
	inputFilePath := aoc.InputArg()
	Main(inputFilePath)
}

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)
	size := 10007
	d := NewDeck(size)
	td := NewTrackingDeck(size, 2019)
	d.Shuffle(lines)
	td.Track(lines)

	fmt.Println(d.cl.ToSlice(d.direction)[:10])
	fmt.Println(d.cl.ToSlice(d.direction)[size-10:])

	fmt.Println("Exercise 1:", d.Position(2019))
	fmt.Println("           ", td.Index)
	fmt.Println("           ", trackFunctions(lines, 2, size))
}

func trackFunctions(lines []string, index, size int) int {
	ops := parse(lines)
	f := function{indexMultiplier: 1}
	for i := len(ops) - 1; i >= 0; i-- {
		// for i := 0; i < len(ops); i++ {
		op := ops[i]
		switch op[0] {
		case cut:
			f = Uncut(f, op[1])
			f = f.Normalize(size)
		case inc:
			f = Uninc(f, size, op[1])
			f = f.Normalize(size)
		case deal:
			f = Undeal(f)
			f = f.Normalize(size)
		}
	}
	fmt.Println(f)
	return f.Value(index, size)
}

func (d *Deck) Shuffle(instructions []string) {
	for _, op := range parse(instructions) {
		switch op[0] {
		case cut:
			d.Cut(op[1])
		case inc:
			d.Increment(op[1])
		case deal:
			d.Deal()
		}
	}
}

type Deck struct {
	size      int
	direction collections.Direction
	cl        *collections.Deque
}

func NewDeck(size int) *Deck {
	cl := collections.NewDeque()
	for v := 0; v < size; v++ {
		cl.Append(v)
	}
	return &Deck{size, collections.Forward, cl}
}

func (d *Deck) String() string {
	return fmt.Sprintf("%v", d.cl)
}

func (d *Deck) Deal() {
	d.direction = d.direction.Opposite()
}

func (d *Deck) Cut(n int) {
	if d.direction == collections.Forward {
		n = -n
	}
	d.cl.Rotate(n)
}

func (d *Deck) Increment(n int) {
	c := make([]int, d.size)

	index := 0
	for i := 0; i < d.size; i++ {
		if d.direction == collections.Forward {
			c[index], _ = d.cl.PopLeft()
		} else {
			c[index], _ = d.cl.Pop()
		}
		index += n
		index %= d.size
	}

	d.direction = collections.Forward
	d.cl = collections.NewDeque()
	for _, v := range c {
		d.cl.Append(v)
	}
}

func (d *Deck) Content() []int {
	return d.cl.ToSlice(d.direction)
}

func (d *Deck) Position(n int) int {
	for i, v := range d.Content() {
		if n == v {
			return i
		}
	}
	return -1
}

type TrackingDeck struct {
	size, Index int
}

func NewTrackingDeck(size, index int) *TrackingDeck {
	return &TrackingDeck{size, index}
}

func (d *TrackingDeck) Track(instructions []string) {
	for _, op := range parse(instructions) {
		switch op[0] {
		case cut:
			d.Cut(op[1])
		case inc:
			d.Increment(op[1])
		case deal:
			d.Deal()
		}
	}
}

func (d *TrackingDeck) Deal() {
	d.Index = d.size - d.Index - 1
}

func (d *TrackingDeck) Cut(n int) {
	d.Index = (d.Index + d.size - n) % d.size
}

func (d *TrackingDeck) Increment(n int) {
	d.Index = d.Index * n % d.size
}

type function struct {
	indexMultiplier int
	constant        int
}

func Deal(f function) function {
	// deal(i) = -i - 1
	return function{
		indexMultiplier: -f.indexMultiplier,
		constant:        -1 - f.constant,
	}
}

func Cut(f function, n int) function {
	// cut(i, n) = i - n
	return function{
		indexMultiplier: f.indexMultiplier,
		constant:        f.constant - n,
	}
}

func Inc(f function, n int) function {
	// inc(i, n) = i * n
	return function{
		indexMultiplier: f.indexMultiplier * n,
		constant:        f.constant * n,
	}
}

func Undeal(f function) function {
	//   deal(i) = -i - 1
	// undeal(i) = -i - 1
	return function{
		indexMultiplier: -f.indexMultiplier,
		constant:        -1 - f.constant,
	}
}

func Uncut(f function, n int) function {
	//   cut(i, n) = i - n
	// uncut(i, n) = i + n
	return function{
		indexMultiplier: f.indexMultiplier,
		constant:        f.constant + n,
	}
}

func Uninc(f function, s, n int) function {
	// s            10007
	// n            7
	// s / n        1429
	// s % n        4
	// n-s%n        3
	// i % n        0    1    2    3    4    5    6  0    1    2
	// i / n        0    0    0    0    0    0    0  1    1    1
	// (n-s%n)*i%n  0    5    3    1    6    4    2  0    5    3
	// result       0 7148 4289 1430 8578 5719 2860  1 7149 4290

	//    inc(i, s, n) = i * n
	//	uninc(i, s, n) = (n * (n-i%n) + i//n + 1)

	mapping := map[int]int{}
	reminder := s % n

	for i := 0; i < n; i++ {
		mapping[(n-reminder)*i%n] = i
	}

	return function{
		indexMultiplier: f.indexMultiplier * n,
		constant:        f.constant * n,
	}
}

func (f function) Apply(of function) function {
	return function{
		indexMultiplier: f.indexMultiplier * of.indexMultiplier,
		constant:        f.indexMultiplier*of.constant + f.constant,
	}
}

func (f function) String() string {
	iSign := ""
	i := f.indexMultiplier
	if i < 0 {
		i *= -1
		iSign = "-"
	}

	constSign := "+"
	constant := f.constant
	if constant < 0 {
		constant *= -1
		constSign = "-"
	}
	return fmt.Sprintf("f(i) = %si*%d %s %d", iSign, i, constSign, constant)
}

func (f function) Normalize(n int) function {
	return function{
		indexMultiplier: f.indexMultiplier % n,
		constant:        f.constant % n,
	}
}

func (f function) Value(i, s int) int {
	f = f.Normalize(s)
	v := ((i*f.indexMultiplier)%s + f.constant) % s
	if v < 0 {
		v = s + v
	}
	return v
}
