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
	size int
}

func NewCircularList() *Deque {
	return &Deque{}
}

func (cl *Deque) Add(v int) {
	cl.size++
	if cl.head == nil {
		cl.head = &node{v, nil, nil}
		cl.head.next = cl.head
		cl.head.prev = cl.head
	} else {
		n := &node{v, nil, nil}
		cl.head.next, cl.head, n.prev, n.next = n, n, cl.head, cl.head.next
		fmt.Println(

		cl.head
	}
}

func (cl *Deque) ToSlice(d Direction) []int {
	content := []int{}
	n := cl.head
	for i := 0; i < cl.size; i++ {
		content = append(content, n.v)
		if d == Forward {
			n = n.next
		} else {
			n = n.prev
		}
	}
	return content
}

func (cl *Deque) LShift() {
	cl.head = cl.head.prev
}

func (cl *Deque) RShift() {
	cl.head = cl.head.next
}

func (cl *Deque) LPop() int {
	v := cl.head.v
	cl.LShift()
	cl.size--
	return v
}

func (cl *Deque) RPop() int {
	v := cl.head.v
	cl.RShift()
	cl.size--
	return v
}

func (cl *Deque) Size() int {
	return cl.size
}

func (cl *Deque) String() string {
	repr := []string{}
	p := cl.head
	fmt.Println(cl.size)
	for i := 0; i < cl.size; i++ {
		repr = append(repr, fmt.Sprintf("%v", p))
		p = p.next
	}

	return strings.Join(repr, ", ")
}
