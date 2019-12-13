package main

import (
	"aoc"
	"aoc/intcode"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	inputFilePath := aoc.InputArg()
	Main(inputFilePath)
}

const (
	empty            = 0
	wall             = 1
	block            = 2
	horizontalPaddle = 3
	ball             = 4
)

const (
	neutral   = 0
	tiltLeft  = -1
	tiltRight = 1
)

const (
	x = 0
	y = 1
)

type cord [2]int

type tile int

var arcade *Arcade

var tString = map[tile]string{
	empty:            " ",
	wall:             "#",
	block:            "▮",
	horizontalPaddle: "—",
	ball:             "•",
}

type Arcade struct {
	cpu  *intcode.Intcode
	grid map[cord]tile
}

func NewArcade(cpu *intcode.Intcode) *Arcade {
	return &Arcade{
		cpu,
		map[cord]tile{},
	}
}

func (a *Arcade) StartGame() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(reader.ReadByte())
	for {
		m := intcode.OutputMode
		for m != intcode.HaltMode {
			x, m := a.cpu.Output()
			if m != intcode.OutputMode {
				break
			}
			y, m := a.cpu.Output()
			if m != intcode.OutputMode {
				panic("Missing y cord!")
			}
			t, m := a.cpu.Output()
			a.grid[cord{x, y}] = tile(t)
		}
	}
}

func (a *Arcade) Draw() {
	m := intcode.OutputMode
	for m != intcode.HaltMode {
		x, m := a.cpu.Output()
		if m != intcode.OutputMode {
			break
		}
		y, m := a.cpu.Output()
		if m != intcode.OutputMode {
			panic("Missing y cord!")
		}
		t, m := a.cpu.Output()
		a.grid[cord{x, y}] = tile(t)
	}
}

func (a *Arcade) String() string {
	var str []string

	maxx := 0
	maxy := 0

	for c := range a.grid {
		if c[x] > maxx {
			maxx = c[x]
		}
		if c[y] > maxy {
			maxy = c[y]
		}
	}

	for y := 0; y <= maxy; y++ {
		line := []string{}
		for x := 0; x <= maxx; x++ {
			line = append(line, tString[a.grid[cord{x, y}]])
		}
		str = append(str, strings.Join(line, ""))
	}

	score := a.grid[cord{-1, 0}]
	str = append(str, fmt.Sprintf("\n SCORE: %d", score))

	return strings.Join(str, "\n")
}

type interactiveInput struct {
	in chan int
}

func (c *interactiveInput) Add(v int) {
}

func (c *interactiveInput) Get() int {
	var ballPosition, paddlePosition cord
	hasBlock := false

	for c, t := range arcade.grid {
		if t == horizontalPaddle {
			paddlePosition = c
		}
		if t == ball {
			ballPosition = c
		}
		if t == block {
			hasBlock = true
		}
	}

	if !hasBlock {
		return neutral
	}

	if ballPosition[x] > paddlePosition[x] {
		return tiltRight
	}
	if ballPosition[x] < paddlePosition[x] {
		return tiltLeft
	}

	return neutral
}

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)
	line := strings.Join(lines, "")

	fmt.Printf("Exercise 1: %d\n", Exercise1(line))

	cpu := intcode.NewIntcodeInput(intcode.NewMemory(line), &chanInput{})
	cpu.SetMemory(0, 2)
	arcade = NewArcade(cpu)
	arcade.StartGame()
}

func Exercise1(line string) int {
	cpu := intcode.NewIntcode(intcode.NewMemory(line))
	arcade := NewArcade(cpu)
	arcade.Draw()

	hm := 0
	for _, t := range arcade.grid {
		if t == block {
			hm++
		}
	}

	return hm
}
