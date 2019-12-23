package main

import (
	"aoc"
	"aoc/intcode"
	"fmt"
)

func main() {
	inputFilePath := aoc.InputArg()
	Main(inputFilePath)
}

type networkInput struct {
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
		return -1
	}
	value := i.values[i.inputPosition]
	fmt.Printf("IN %2d: %d\n", i.address, value)
	i.inputPosition++
	return value
}

func run(i int, cpu *intcode.Intcode, inputs map[int]*networkInput, finished chan bool) {
	for {
		destination, m := cpu.Output()
		if m == intcode.HaltMode {
			fmt.Printf("%2d: STOP a %d\n", i, destination)
			break
		}
		x, m := cpu.Output()
		if m == intcode.HaltMode {
			fmt.Printf("%2d: STOP b %d\n", i, destination)
			break
		}
		y, m := cpu.Output()
		if m == intcode.HaltMode {
			fmt.Printf("%2d: STOP c %d\n", i, destination)
			break
		}
		fmt.Printf("%2d: > %2d (%6d %6d)\n", i, destination, x, y)
		if destination == 255 {
			panic("!")
		}
		inputs[destination].AddPacket(x, y)
	}
	finished <- true
}

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)
	cpus := map[int]*intcode.Intcode{}
	inputs := map[int]*networkInput{}
	for i := 0; i < 50; i++ {
		inputs[i] = &networkInput{address: i}
		inputs[i].Add(i)
		cpus[i] = intcode.NewIntcodeInput(intcode.NewMemory(lines[0]), inputs[i])
	}

	finished := make(chan bool)

	for i, cpu := range cpus {
		fmt.Println("Running", i)
		go run(i, cpu, inputs, finished)
	}

	for _ = range cpus {
		<-finished
	}
}
