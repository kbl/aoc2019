package main

import (
	"aoc"
	"fmt"
	"log"
	"strconv"
	"strings"
)

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
	// too high 9209
	fmt.Println("Exercise 1:", d.Position(2019))
	fmt.Println("Exercise 1:", td.Index)
	size := 101741582076661
	noShufles := 119315717514047

	td.Track(lines)
	fmt.Println(size)
	fmt.Println(noShufles)
	fmt.Println("Exercise 2:", td.Index)

	index := 2020
	i := 0

	for {
		td = NewTrackingDeck(size, index)
		td.Track(lines)
		index = td.Index
		if i%10000 == 0 {
			fmt.Printf("%6d: %d\n", i, td.Index)
		}
		if index == 2020 {
			fmt.Printf("%6d: %d\n", i, td.Index)
			break
		}
		i++
	}
}

func (d *Deck) Shuffle(instructions []string) {
	for _, l := range instructions {
		t := strings.Split(l, " ")
		if t[0] == "cut" {
			v, err := strconv.Atoi(t[1])
			if err != nil {
				log.Fatal(err)
			}
			d.Cut(v)
		} else if t[1] == "with" {
			v, err := strconv.Atoi(t[3])
			if err != nil {
				log.Fatal(err)
			}
			d.Increment(v)
		} else if t[1] == "into" {
			d.Deal()
		} else {
			fmt.Println(l)
			panic("unknown")
		}
	}
}

type node struct {
	v          int
	prev, next *node
}

const (
	forward  = 0
	backward = 1
)

type Deck struct {
	size, direction int
	pointer         *node
}

func NewDeck(size int) *Deck {
	first := &node{}
	prev := first
	for v := 1; v < size; v++ {
		n := &node{v: v, prev: prev, next: nil}
		prev.next = n
		prev = n
	}
	prev.next = first
	first.prev = prev
	return &Deck{size, forward, first}
}

func (d *Deck) String() string {
	repr := []string{}
	for _, v := range d.Content() {
		repr = append(repr, strconv.Itoa(v))
	}
	return strings.Join(repr, ", ")
}

func (d *Deck) Deal() {
	if d.direction == forward {
		d.direction = backward
	} else {
		d.direction = forward
	}
	d.Cut(1)
}

func (d *Deck) Cut(n int) {
	direction := d.direction
	if n < 0 {
		n = -n
		direction = forward
		if d.direction == forward {
			direction = backward
		}
	}
	for i := 0; i < n; i++ {
		if direction == forward {
			d.pointer = d.pointer.next
		} else {
			d.pointer = d.pointer.prev
		}
	}
}

func (d *Deck) Increment(n int) {
	c := make([]int, d.size)

	index := 0
	p := d.pointer
	for i := 0; i < d.size; i++ {
		c[index] = p.v
		index += n
		index %= d.size
		if d.direction == forward {
			p = p.next
		} else {
			p = p.prev
		}
	}

	d.direction = forward
	node := d.pointer
	for _, v := range c {
		node.v = v
		node = node.next
	}
}

func (d *Deck) Position(n int) int {
	node := d.pointer
	for i := 0; i < d.size; i++ {
		if n == node.v {
			return i
		}
		if d.direction == forward {
			node = node.next
		} else {
			node = node.prev
		}
	}

	return -1
}

func (d *Deck) Content() []int {
	content := []int{}
	p := d.pointer
	for i := 0; i < d.size; i++ {
		content = append(content, p.v)
		if d.direction == forward {
			p = p.next
		} else {
			p = p.prev
		}
	}
	return content
}

type TrackingDeck struct {
	size, Index int
}

func NewTrackingDeck(size, index int) *TrackingDeck {
	return &TrackingDeck{size, index}
}

func (d *TrackingDeck) Track(instructions []string) {
	for _, l := range instructions {
		t := strings.Split(l, " ")
		if t[0] == "cut" {
			v, err := strconv.Atoi(t[1])
			if err != nil {
				log.Fatal(err)
			}
			d.Cut(v)
		} else if t[1] == "with" {
			v, err := strconv.Atoi(t[3])
			if err != nil {
				log.Fatal(err)
			}
			d.Increment(v)
		} else if t[1] == "into" {
			d.Deal()
		} else {
			fmt.Println(l)
			panic("unknown")
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
	int size
	int mSize
	int index
	int mIndex
	int constant
}

func NewFunction(size, index int) *Function {
	return &Function{size, index}
}
