package collections

import (
	"fmt"
	"strings"
)

type Direction int

const (
	Forward  = 0
	Backward = 1
)

func (d Direction) Opposite() Direction {
	if d == Forward {
		return Backward
	}
	return Forward
}

type node struct {
	v          int
	prev, next *node
}

type Deque struct {
	head *node
	tail *node
	size int
}

func NewDeque() *Deque {
	return &Deque{}
}

func (d *Deque) Append(v int) {
	d.size++
	if d.head == nil {
		d.head = &node{v, nil, nil}
		d.tail = d.head
	} else {
		// n1 <-> n2 -> nil
		n1 := d.head
		n2 := &node{
			v:    v,
			prev: n1,
		}
		n1.next = n2
		d.head = n2
	}
}

func (d *Deque) Pop() (int, bool) {
	if d.head != nil {
		d.size--
		v := d.head.v
		d.head = d.head.prev
		if d.head != nil {
			d.head.next = nil
		}
		return v, true
	}
	return 0, false
}

func (d *Deque) ToSlice(dir Direction) []int {
	content := []int{}

	if dir == Forward {
		n := d.tail
		for n != nil {
			content = append(content, n.v)
			n = n.next
		}
	} else {
		n := d.head
		for n != nil {
			content = append(content, n.v)
			n = n.prev
		}
	}

	return content
}

func (d *Deque) Rotate(x int) {
	old_head := d.head
	old_tail := d.tail
	if x > 0 {
		n := d.head

		for i := 0; i < x; i++ {
			n = n.prev
		}

		d.head = n
		d.tail = n.next
	} else {
		n := d.tail

		for i := 0; i > x; i-- {
			n = n.next
		}

		d.tail = n
		d.head = n.prev
	}
	old_tail.prev = old_head
	old_head.next = old_tail

	d.tail.prev = nil
	d.head.next = nil
}

func (d *Deque) Size() int {
	return d.size
}

func (d *Deque) String() string {
	repr := []string{}
	n := d.tail
	for n != nil {
		repr = append(repr, fmt.Sprintf("%d", n.v))
		n = n.next
	}

	return strings.Join(repr, ", ")
}
