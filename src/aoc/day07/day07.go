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

	ics := NewIntcodeSequence(memory, []int{0, 1, 2, 3, 4})
	fmt.Println("Exercise 1: ", ics.Run())

	ics = NewIntcodeSequence(memory, []int{5, 6, 7, 8, 9})
	fmt.Println("Exercise 2: ", ics.Run())
}

func parseMemory(line string) []int {
	strNumbers := strings.Split(line, ",")
	var memory []int
	for _, sn := range strNumbers {
		m, err := strconv.Atoi(sn)
		if err != nil {
			log.Fatal(err)
		}
		memory = append(memory, m)
	}
	return memory
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

type IntcodeSequence struct {
	memory, phaseOptions []int
}

func NewIntcodeSequence(memory, phaseOptions []int) *IntcodeSequence {
	return &IntcodeSequence{
		memory,
		phaseOptions,
	}
}

// https://www.golangprograms.com/golang-program-to-generate-slice-permutations-of-number-entered-by-user.html
func (is *IntcodeSequence) phaseSettings() (settings [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			settings = append(settings, append([]int{}, a...))
		} else {
			for i := k; i < len(is.phaseOptions); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(is.phaseOptions, 0)

	return settings
}

func (is *IntcodeSequence) Run() int {
	biggestOutput := 0

	for _, phaseSetting := range is.phaseSettings() {
		intcodes := []*Intcode{}

		for _, ps := range phaseSetting {
			ic := NewIntcode(is.memory)
			ic.AddInput(ps)
			intcodes = append(intcodes, ic)
		}

		previousOutput := 0
		m := haltMode

		for {
			for _, ic := range intcodes {
				ic.AddInput(previousOutput)
				previousOutput, m = ic.Output()
			}

			if m == haltMode {
				break
			}
		}

		if previousOutput > biggestOutput {
			biggestOutput = previousOutput
		}
	}

	return biggestOutput
}

type mode int

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
	instructionPointer int
	memory             []int
	in                 *Input
	out                int
}

func NewIntcode(memory []int) *Intcode {
	return &Intcode{
		memory: memory,
		in:     NewInput(),
	}
}

var haltMode mode = 0
var outputMode mode = 1

func (i *Intcode) AddInput(value int) {
	i.in.Add(value)
}

func (i *Intcode) Output() (int, mode) {
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
			return i.out, haltMode
		case input:
			i.input()
		case output:
			return i.output(modes), outputMode
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

func (i *Intcode) output(modes *modes) int {
	outputValue := i.value(0, modes, i.instructionPointer+1)
	i.out = outputValue
	i.instructionPointer += 2
	return outputValue
}

func (i *Intcode) jumpIfTrue(modes *modes) {
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
