package main

import (
	"aoc"
	"aoc/collections"
	"fmt"
	"log"
	"math/big"
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
	d.Shuffle(lines)

	fmt.Println(d.cl.ToSlice(d.direction)[:10])

	f := toFunction(lines, size)
	fmt.Println("Exercise 1:", f.Value(2019, size))
	fmt.Println("           ", f)

	size = 119315717514047
	reps := 101741582076661
	f = toFunction(lines, size)
	fmt.Println("Exercise 2:", exercise2(f, size, reps))

	// f(i) = i*44064224361553 + 47730710794979
	//
	// wolframalpha to the rescue ;)
	// 2020 = i*44064224361553 + 47730710794979 mod 119315717514047
	// 78613970589919
}

func exercise2(f function, size, reps int) function {
	powers := []int{1}
	powerToF := map[int]function{
		1: f,
	}

	for i := 2; i <= reps; i *= 2 {
		f = f.Apply(f)
		f = f.Normalize(size)
		powerToF[i] = f
		powers = append(powers, i)
	}

	finalFunction := function{indexMultiplier: big.NewInt(1), constant: big.NewInt(0)}

	for i := len(powers) - 1; i >= 0; i-- {
		p := powers[i]
		if reps >= p {
			reps -= p
			finalFunction = finalFunction.Apply(powerToF[p])
			finalFunction = finalFunction.Normalize(size)
		}
	}

	return finalFunction
}

func toFunction(lines []string, size int) function {
	ops := parse(lines)
	f := function{indexMultiplier: big.NewInt(1), constant: big.NewInt(0)}
	for i := 0; i < len(ops); i++ {
		op := ops[i]
		switch op[0] {
		case cut:
			f = Cut(f, op[1])
			f = f.Normalize(size)
		case inc:
			f = Inc(f, op[1])
			f = f.Normalize(size)
		case deal:
			f = Deal(f)
			f = f.Normalize(size)
		}
	}
	return f
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

type function struct {
	indexMultiplier *big.Int
	constant        *big.Int
}

func Deal(f function) function {
	// deal(i) = -i - 1
	return function{
		indexMultiplier: big.NewInt(0).Mul(f.indexMultiplier, big.NewInt(-1)),
		constant:        big.NewInt(0).Sub(big.NewInt(-1), f.constant),
	}
}

func Cut(f function, n int) function {
	// cut(i, n) = i - n
	return function{
		indexMultiplier: f.indexMultiplier,
		constant:        big.NewInt(0).Sub(f.constant, big.NewInt(int64(n))),
	}
}

func Inc(f function, n int) function {
	// inc(i, n) = i * n
	nn := big.NewInt(int64(n))
	return function{
		indexMultiplier: big.NewInt(0).Mul(f.indexMultiplier, nn),
		constant:        big.NewInt(0).Mul(f.constant, nn),
	}
}

func (f function) Apply(of function) function {
	return function{
		indexMultiplier: big.NewInt(0).Mul(f.indexMultiplier, of.indexMultiplier),
		constant:        big.NewInt(0).Add(f.constant, big.NewInt(0).Mul(of.constant, f.indexMultiplier)),
	}
}

func (f function) String() string {
	iSign := ""
	i := f.indexMultiplier
	if i.Sign() < 0 {
		i = big.NewInt(1).Mul(big.NewInt(-1), i)
		iSign = "-"
	}

	constSign := "+"
	constant := f.constant
	if constant.Sign() < 0 {
		constant = big.NewInt(1).Mul(big.NewInt(-1), constant)
		constSign = "-"
	}
	return fmt.Sprintf("f(i) = %si*%d %s %d", iSign, i, constSign, constant)
}

func (f function) Normalize(n int) function {
	nn := big.NewInt(int64(n))
	return function{
		indexMultiplier: big.NewInt(0).Mod(f.indexMultiplier, nn),
		constant:        big.NewInt(0).Mod(f.constant, nn),
	}
}

func (f function) Value(i, s int) int {
	f = f.Normalize(s)
	im := f.indexMultiplier.Int64()
	c := f.constant.Int64()
	v := ((int64(i)*im)%int64(s) + c) % int64(s)
	if v < 0 {
		v = int64(s) + v
	}
	return int(v)
}
