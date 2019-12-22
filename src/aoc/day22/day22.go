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
	d.Shuffle(lines)
	// too high 9209
	fmt.Println("Exercise 1:", d.Position(2019))
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
	return d.Content()[n]
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
