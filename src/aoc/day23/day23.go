package main

import (
	"aoc"
	"aoc/intcode"
	"fmt"
	"time"
)

func main() {
	inputFilePath := aoc.InputArg()
	Main(inputFilePath)
}

type networkInput struct {
	isIdle        bool
	address       int
	values        []int
	inputPosition int
}

func (i *networkInput) Add(value int) {
	i.values = append(i.values, value)
}

func (i *networkInput) AddPacket(x, y int) {
	temp := append(i.values, x)
	temp = append(temp, y)
	i.values = temp
}

func (i *networkInput) Get() int {
	if i.inputPosition >= len(i.values) {
		i.isIdle = true
		return -1
	}
	i.isIdle = false
	value := i.values[i.inputPosition]
	i.inputPosition++
	return value
}

type nat struct {
	x, y     int
	previous int
	inputs   map[int]*networkInput
}

func (n *nat) monitor() {
	first := true
	for {
		time.Sleep(10000)
		idle := true
		for _, input := range n.inputs {
			idle = idle && input.isIdle
		}

		if idle && !first && n.x+n.y != 0 {
			fmt.Printf("NAT: %d %d\n", n.x, n.y)
			n.inputs[0].AddPacket(n.x, n.y)
			if !first && n.y == n.previous {
				fmt.Println(n.y)
				panic("double!")
			}
			first = false
			n.previous = n.y
		}
	}
}

func run(i int, cpu *intcode.Intcode, inputs map[int]*networkInput, finished chan bool, n *nat) {
	for {
		destination, m := cpu.Output()
		if m == intcode.HaltMode {
			fmt.Printf("%3d: STOP a %d\n", i, destination)
			break
		}
		x, m := cpu.Output()
		if m == intcode.HaltMode {
			fmt.Printf("%3d: STOP b %d\n", i, destination)
			break
		}
		y, m := cpu.Output()
		if m == intcode.HaltMode {
			fmt.Printf("%3d: STOP c %d\n", i, destination)
			break
		}
		if destination == 255 {
			fmt.Printf("%3d > NAT (%6d %6d)\n", i, x, y)
			fmt.Println(n)
			n.x, n.y = x, y
		} else {
			fmt.Printf("%2d: > %3d (%6d %6d)\n", i, destination, x, y)
			inputs[destination].AddPacket(x, y)
		}
	}
	finished <- true
}

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)
	cpus := map[int]*intcode.Intcode{}
	inputs := map[int]*networkInput{}
	n := &nat{inputs: inputs}
	for i := 0; i < 50; i++ {
		inputs[i] = &networkInput{address: i}
		inputs[i].Add(i)
		cpus[i] = intcode.NewIntcodeInput(intcode.NewMemory(lines[0]), inputs[i])
	}

	finished := make(chan bool)

	for i, cpu := range cpus {
		go run(i, cpu, inputs, finished, n)
	}
	go n.monitor()

	<-finished
}
