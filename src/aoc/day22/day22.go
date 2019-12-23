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
	d := NewDeck(10007)
	td := NewTrackingDeck(10007, 2019)
	d.Shuffle(lines)
	td.Track(lines)

	fmt.Println("Exercise 1:", d.Position(2019))
	fmt.Println("Exercise 1:", td.Index)
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

const (
	forward  = 0
	backward = 1
)

type Deck struct {
	size      int
	direction collections.Direction
	cl        *collections.CircularList
}

func NewDeck(size int) *Deck {
	cl := collections.NewCircularList()
	for v := 0; v < size; v++ {
		cl.Add(v)
	}
	cl.RShift()
	return &Deck{size, forward, cl}
}

func (d *Deck) String() string {
	repr := []string{}
	for _, v := range d.cl.ToSlice(d.direction) {
		repr = append(repr, strconv.Itoa(v))
	}
	return strings.Join(repr, ", ")
}

func (d *Deck) Deal() {
	fmt.Println(1, d)
	fmt.Println(d.cl)
	d.direction = d.direction.Opposite()
	fmt.Println(2, d)
	fmt.Println(d.cl)
	d.Cut(1)
	fmt.Println(3, d)
	fmt.Println(d.cl)
	fmt.Println()
}

func (d *Deck) Cut(n int) {
	direction := d.direction
	if n < 0 {
		n = -n
		direction = d.direction.Opposite()
	}
	for i := 0; i < n; i++ {
		if direction == collections.Forward {
			d.cl.RShift()
		} else {
			d.cl.LShift()
		}
	}
}

func (d *Deck) Increment(n int) {
	c := make([]int, d.size)

	index := 0
	for i := 0; i < d.size; i++ {
		if d.direction == collections.Forward {
			c[index] = d.cl.RPop()
		} else {
			c[index] = d.cl.LPop()
		}
		index += n
		index %= d.size
	}

	d.direction = collections.Forward
	d.cl = collections.NewCircularList()
	for _, v := range c {
		d.cl.Add(v)
	}
	d.cl.RShift()
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

//	deal() = s - i - 1
//	cut(n) = (i + s - n) % s
//	inc(n) = i * n

type Function struct {
	sizeMultiplier  int
	indexMultiplier int
	constant        int
}

var Deal = func() Function {
	return Function{
		sizeMultiplier:  1,
		indexMultiplier: -1,
		constant:        -1,
	}
}

var Cut = func(n int) Function {
	return Function{
		sizeMultiplier:  1,
		indexMultiplier: 1,
		constant:        -n,
	}
}

var Inc = func(n int) Function {
	return Function{
		sizeMultiplier:  0,
		indexMultiplier: n,
		constant:        0,
	}
}

func (f Function) String() string {
	return fmt.Sprintf("f(i, s) = i*%d + s*%d + %d", f.indexMultiplier, f.sizeMultiplier, f.constant)
}

func (f Function) Apply(of Function) Function {
	return Function{
		indexMultiplier: f.indexMultiplier + of.indexMultiplier,
		sizeMultiplier:  f.sizeMultiplier + of.sizeMultiplier,
		constant:        f.constant + of.constant,
	}
}

func (f Function) Normalize(n int) Function {
	if f.indexMultiplier < 0 || f.sizeMultiplier < 0 || f.constant < 0 {
		panic("Can't normalize negative values!")
	}
	return Function{
		indexMultiplier: f.indexMultiplier % n,
		sizeMultiplier:  f.sizeMultiplier % n,
		constant:        f.constant % n,
	}
}

func (f Function) Value(i, s int) int {
	f = f.Normalize(s)
	return ((i*f.indexMultiplier)%s + (s*f.sizeMultiplier)%s + f.constant) % s
}
