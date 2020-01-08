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
			f = Uninc(f, op[1])
			f = f.Normalize(size)
		case deal:
			f = Undeal(f)
			f = f.Normalize(size)
		}
	}
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
	sizeMultiplier  int
	indexMultiplier int
	constant        int
}

func Deal(f function) function {
	//	deal(i, s) = s - i - 1
	return function{
		sizeMultiplier:  1 - f.sizeMultiplier,
		indexMultiplier: -f.indexMultiplier,
		constant:        -1 - f.constant,
	}
}

func Cut(f function, n int) function {
	//	cut(i, s, n) = i + s - n
	return function{
		sizeMultiplier:  1 + f.sizeMultiplier,
		indexMultiplier: f.indexMultiplier,
		constant:        f.constant - n,
	}
}

func Inc(f function, n int) function {
	//	inc(i, s, n) = i * n
	return function{
		sizeMultiplier:  f.sizeMultiplier * n,
		indexMultiplier: f.indexMultiplier * n,
		constant:        f.constant * n,
	}
}

func Undeal(f function) function {
	//   deal(i, s) = s - i - 1
	// undeal(i, s) = s - i - 1
	return function{
		sizeMultiplier:  1 - f.sizeMultiplier,
		indexMultiplier: -f.indexMultiplier,
		constant:        -1 - f.constant,
	}
}

func Uncut(f function, n int) function {
	//   cut(i, s, n) = i + s - n
	// uncut(i, s, n) = i + s + n
	return function{
		sizeMultiplier:  1 + f.sizeMultiplier,
		indexMultiplier: f.indexMultiplier,
		constant:        f.constant + n,
	}
}

// s = 10
// n = 9
// 0 -> 0
// 1 -> 9
// 2 -> 8
// 3 -> 7
// 4 -> 6
// 5 -> 5
// 6 -> 4
// 7 -> 3
// 8 -> 2
// 9 -> 1
//
// n = 7
// 0 -> 0
// 1 -> 3
// 2 -> 6
// 3 -> 9
// 4 -> 2
// 5 -> 5
// 6 -> 8
// 7 -> 1
// 8 -> 4
// 9 -> 7
//
// n = 1
// 0 -> 0
// 1 -> 1
// 2 -> 2
// 3 -> 3
// 4 -> 4
// 5 -> 5
// 6 -> 6
// 7 -> 7
// 8 -> 8
// 9 -> 9

// s = 10007
// n = 7
// 0 -> 0
// 1 -> 7
// 2 -> 14
// 3 -> 21
// …

func Uninc(f function, n int) function {
	//	uninc(i, s, n) = s - i * n
	return function{
		sizeMultiplier:  f.sizeMultiplier * n,
		indexMultiplier: f.indexMultiplier * n,
		constant:        f.constant * n,
	}
}

func (f function) Apply(of function) function {
	return function{
		sizeMultiplier:  f.indexMultiplier*of.sizeMultiplier + f.sizeMultiplier,
		indexMultiplier: f.indexMultiplier * of.indexMultiplier,
		constant:        f.indexMultiplier*of.constant + f.constant,
	}
}

func (f function) String() string {
	return fmt.Sprintf("f(i, s) = i*%d + s*%d + %d", f.indexMultiplier, f.sizeMultiplier, f.constant)
}

func (f function) Normalize(n int) function {
	return function{
		indexMultiplier: f.indexMultiplier % n,
		sizeMultiplier:  f.sizeMultiplier % n,
		constant:        f.constant % n,
	}
}

func (f function) Value(i, s int) int {
	f = f.Normalize(s)
	v := ((i*f.indexMultiplier)%s + (s*f.sizeMultiplier)%s + f.constant) % s
	if v < 0 {
		v = s + v
	}
	return v
}
