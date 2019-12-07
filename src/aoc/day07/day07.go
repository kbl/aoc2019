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
	memory := parseMemory(lines[0])

	fmt.Printf("Exercise 1: ")
	exercise1(memory, 1)
	// fmt.Printf("Exercise 2: ")
	// exercise1(memory, 5)
}

func parseMemory(line string) *[]int {
	strNumbers := strings.Split(line, ",")
	var memory []int
	for _, sn := range strNumbers {
		m, err := strconv.Atoi(sn)
		if err != nil {
			log.Fatal(err)
		}
		memory = append(memory, m)
	}
	return &memory
}

type Input struct {
	values        []int
	inputPosition int
}

func NewInput() *Input {
	return &Input{
		[]int{},
		0,
	}
}

func (i *Input) Add(value int) {
	i.values = append(i.values, value)
}

func (i *Input) Get() int {
	value := i.values[i.inputPosition]
	i.inputPosition++
	return value
}

func permutation(xs []int) (permuts [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
}

func exercise1(memoryPtr *[]int, inputValue int) {
	memory := *memoryPtr
	biggestOutput := 0

	ps := permutation([]int{0, 1, 2, 3, 4})

	for _, p := range ps {
		a := p[0]
		b := p[1]
		c := p[2]
		d := p[3]
		e := p[4]

		memCopy := make([]int, len(memory))
		copy(memCopy, memory[:])
		intcodeA := NewIntcode(memCopy)
		intcodeA.AddInput(a)
		intcodeA.AddInput(0)

		memCopy = make([]int, len(memory))
		copy(memCopy, memory[:])
		intcodeB := NewIntcode(memCopy)
		intcodeB.AddInput(b)

		memCopy = make([]int, len(memory))
		copy(memCopy, memory[:])
		intcodeC := NewIntcode(memCopy)
		intcodeC.AddInput(c)

		memCopy = make([]int, len(memory))
		copy(memCopy, memory[:])
		intcodeD := NewIntcode(memCopy)
		intcodeD.AddInput(d)

		memCopy = make([]int, len(memory))
		copy(memCopy, memory[:])
		intcodeE := NewIntcode(memCopy)
		intcodeE.AddInput(e)

		for {
			outputA := intcodeA.Output()

			intcodeB.AddInput(outputA)
			outputB := intcodeB.Output()

			intcodeC.AddInput(outputB)
			outputC := intcodeC.Output()

			intcodeD.AddInput(outputC)
			outputD := intcodeD.Output()

			intcodeE.AddInput(outputD)
			outputE := intcodeE.Output()

			fmt.Printf("%d%d%d%d%d: %d\n", a, b, c, d, e, outputE)
		}

		if outputE > biggestOutput {
			biggestOutput = outputE
		}
	}

	fmt.Println(biggestOutput)
}

const (
	add         = 1
	multiply    = 2
	input       = 3
	output      = 4
	jumpIfTrue  = 5
	jumpIfFalse = 6
	lessThan    = 7
	equals      = 8
	halt        = 99
)

const (
	position  = 0
	immediate = 1
)

type Intcode struct {
	instructionPointer, DiagnosticCode int
	memory                             []int
	in                                 *Input
}

func NewIntcode(memory []int) *Intcode {
	s := 0
	for _, m := range memory {
		s += m
	}
	// fmt.Println("sum", s)
	instructionPointer := 0
	diagnosticCode := -1
	return &Intcode{
		instructionPointer,
		diagnosticCode,
		memory,
		NewInput(),
	}
}

func (i *Intcode) AddInput(value int) {
	i.in.Add(value)
}

func (i *Intcode) Output() int {
	for {
		opcode := i.memory[i.instructionPointer] % 100
		modes := newModes(i.memory[i.instructionPointer] / 100)
		switch opcode {
		case add:
			i.add(modes)
		case multiply:
			i.multiply(modes)
		case halt:
			i.halt()
			return i.memory[0]
		case input:
			i.input()
		case output:
			i.output(modes)
			return i.DiagnosticCode
		case jumpIfTrue:
			i.jumpIfTrue(modes)
		case jumpIfFalse:
			i.jumpIfFalse(modes)
		case lessThan:
			i.lessThan(modes)
		case equals:
			i.equals(modes)
		default:
			log.Fatalf("Unknown opcode %d\n", opcode)
		}
	}
}

func (i *Intcode) Execute() int {
	for {
		opcode := i.memory[i.instructionPointer] % 100
		modes := newModes(i.memory[i.instructionPointer] / 100)
		switch opcode {
		case add:
			i.add(modes)
		case multiply:
			i.multiply(modes)
		case halt:
			i.halt()
			return i.memory[0]
		case input:
			i.input()
		case output:
			i.output(modes)
		case jumpIfTrue:
			i.jumpIfTrue(modes)
		case jumpIfFalse:
			i.jumpIfFalse(modes)
		case lessThan:
			i.lessThan(modes)
		case equals:
			i.equals(modes)
		default:
			log.Fatalf("Unknown opcode %d\n", opcode)
		}
	}
}

type modes struct {
	modes []int
}

func newModes(intModes int) *modes {
	buffer := []int{}
	for intModes > 0 {
		buffer = append(buffer, intModes%10)
		intModes /= 10
	}
	return &modes{buffer}
}

func (m *modes) getMode(index int) int {
	if index >= len(m.modes) {
		return 0
	}
	return m.modes[index]
}

func (i *Intcode) add(modes *modes) {
	param1Value := i.value(0, modes, i.instructionPointer+1)
	param2Value := i.value(1, modes, i.instructionPointer+2)
	outputIndex := i.memory[i.instructionPointer+3]
	i.memory[outputIndex] = param1Value + param2Value
	i.instructionPointer += 4
}

func (i *Intcode) multiply(modes *modes) {
	param1Value := i.value(0, modes, i.instructionPointer+1)
	param2Value := i.value(1, modes, i.instructionPointer+2)
	outputIndex := i.memory[i.instructionPointer+3]
	i.memory[outputIndex] = param1Value * param2Value
	i.instructionPointer += 4
}

func (i *Intcode) halt() {
	i.instructionPointer += 1
}

func (i *Intcode) input() {
	inputIndex := i.memory[i.instructionPointer+1]
	inputValue := i.in.Get()
	i.memory[inputIndex] = inputValue
	i.instructionPointer += 2
}

func (i *Intcode) output(modes *modes) {
	outputValue := i.value(0, modes, i.instructionPointer+1)
	// fmt.Printf("Diagnostic code: %d\n", outputValue)
	i.DiagnosticCode = outputValue
	i.instructionPointer += 2
}

func (i *Intcode) jumpIfTrue(modes *modes) {
	// jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
	param1Value := i.value(0, modes, i.instructionPointer+1)
	param2Value := i.value(1, modes, i.instructionPointer+2)
	if param1Value != 0 {
		i.instructionPointer = param2Value
	} else {
		i.instructionPointer += 3
	}
}

func (i *Intcode) jumpIfFalse(modes *modes) {
	// jump-if-false: if the first parameter is zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
	param1Value := i.value(0, modes, i.instructionPointer+1)
	param2Value := i.value(1, modes, i.instructionPointer+2)
	if param1Value == 0 {
		i.instructionPointer = param2Value
	} else {
		i.instructionPointer += 3
	}
}

func (i *Intcode) lessThan(modes *modes) {
	// less than: if the first parameter is less than the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
	param1Value := i.value(0, modes, i.instructionPointer+1)
	param2Value := i.value(1, modes, i.instructionPointer+2)
	outputIndex := i.memory[i.instructionPointer+3]
	if param1Value < param2Value {
		i.memory[outputIndex] = 1
	} else {
		i.memory[outputIndex] = 0
	}
	i.instructionPointer += 4
}

func (i *Intcode) equals(modes *modes) {
	// equals: if the first parameter is equal to the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
	param1Value := i.value(0, modes, i.instructionPointer+1)
	param2Value := i.value(1, modes, i.instructionPointer+2)
	outputIndex := i.memory[i.instructionPointer+3]
	if param1Value == param2Value {
		i.memory[outputIndex] = 1
	} else {
		i.memory[outputIndex] = 0
	}
	i.instructionPointer += 4
}

func (i *Intcode) value(paramIndex int, modes *modes, instructionPointer int) int {
	switch modes.getMode(paramIndex) {
	case position:
		paramIndex := i.memory[instructionPointer]
		return i.memory[paramIndex]
	case immediate:
		return i.memory[instructionPointer]
	}
	log.Fatalf("%d mode unknown!\n", modes.getMode(paramIndex))
	return -1
}
