package main

import (
	"aoc"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const debug = false

func main() {
	inputFilePath := aoc.InputArg()
	Main(inputFilePath)
}

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)
	memory := NewMemory(lines[0])

	ic := NewIntcode(memory)
	ic.AddInput(1)
	fmt.Println("Exercise 1: ", ic.Execute())
	ic = NewIntcode(memory)
	ic.AddInput(2)
	fmt.Println("Exercise 2: ", ic.Execute())
}

type Memory struct {
	memory map[int]int
}

func NewMemory(line string) *Memory {
	strNumbers := strings.Split(line, ",")
	memory := map[int]int{}
	for i, sn := range strNumbers {
		m, err := strconv.Atoi(sn)
		if err != nil {
			log.Fatal(err)
		}
		memory[i] = m
	}
	return &Memory{memory}
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

type exitMode int
type paramMode int

const (
	add         = 1
	multiply    = 2
	input       = 3
	output      = 4
	jumpIfTrue  = 5
	jumpIfFalse = 6
	lessThan    = 7
	equals      = 8
	rbOffset    = 9
	halt        = 99
)

const (
	position  = 0
	immediate = 1
	relative  = 2
)

type Intcode struct {
	instructionPointer, relativeBase int
	memory                           *Memory
	in                               *Input
	out                              int
}

func NewIntcode(memory *Memory) *Intcode {
	return &Intcode{
		memory: memory,
		in:     NewInput(),
	}
}

var haltMode exitMode = 0
var outputMode exitMode = 1

func (i *Intcode) AddInput(value int) {
	i.in.Add(value)
}

func (i *Intcode) Execute() int {
	for {
		v, m := i.Output()
		if m == haltMode {
			return v
		}
	}
}

func (i *Intcode) mget(index int) int {
	if v, ok := i.memory.memory[index]; ok {
		return v
	}
	i.memory.memory[index] = 0
	return 0
}

func (i *Intcode) mput(index, value int) {
	if index < 0 {
		panic(index)
	}
	i.memory.memory[index] = value
}

func (i *Intcode) Output() (int, exitMode) {
	for {
		opcode := i.mget(i.instructionPointer) % 100
		modes := newModes(i.mget(i.instructionPointer) / 100)
		switch opcode {
		case add:
			i.add(modes)
		case multiply:
			i.multiply(modes)
		case halt:
			i.halt()
			return i.out, haltMode
		case input:
			i.input(modes)
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
		case rbOffset:
			i.rbOffset(modes)
		default:
			log.Fatalf("Unknown opcode %d\n", opcode)
		}
	}
}

type modes struct {
	modes []paramMode
}

func newModes(intModes int) *modes {
	buffer := []paramMode{}
	for intModes > 0 {
		buffer = append(buffer, paramMode(intModes%10))
		intModes /= 10
	}
	return &modes{buffer}
}

func (m *modes) getMode(index int) paramMode {
	if index >= len(m.modes) {
		return 0
	}
	return m.modes[index]
}

func (m *modes) getModeS(index int) string {
	switch m.getMode(index) {
	case position:
		return "p"
	case immediate:
		return "i"
	case relative:
		return "r"
	}
	panic(fmt.Sprintf("Uknown mode %d!", m.getMode(index)))
}

func (i *Intcode) add(modes *modes) {
	param1Value := i.value(0, modes, i.instructionPointer+1)
	param2Value := i.value(1, modes, i.instructionPointer+2)
	outputIndex := i.valueIndex(2, modes, i.instructionPointer+3)

	if debug {
		fmt.Printf("\nadd: %d %d %d %d\n", i.mget(i.instructionPointer), i.mget(i.instructionPointer+1), i.mget(i.instructionPointer+2), i.mget(i.instructionPointer+3))
		fmt.Printf(" ip: %d\n", i.instructionPointer)
		fmt.Printf(" rb: %d\n", i.relativeBase)
		fmt.Printf("%8d[%s] -> %d\n", i.mget(i.instructionPointer+1), modes.getModeS(0), param1Value)
		fmt.Printf("%8d[%s] -> %d\n", i.mget(i.instructionPointer+2), modes.getModeS(1), param2Value)
		fmt.Printf("%8d[%s] -> m[%d] = %d\n", i.mget(i.instructionPointer+3), modes.getModeS(2), outputIndex, param1Value+param2Value)
	}

	i.mput(outputIndex, param1Value+param2Value)
	i.instructionPointer += 4
}

func (i *Intcode) multiply(modes *modes) {
	param1Value := i.value(0, modes, i.instructionPointer+1)
	param2Value := i.value(1, modes, i.instructionPointer+2)
	outputIndex := i.valueIndex(2, modes, i.instructionPointer+3)

	if debug {
		fmt.Printf("\nmul: %d %d %d %d\n", i.mget(i.instructionPointer), i.mget(i.instructionPointer+1), i.mget(i.instructionPointer+2), i.mget(i.instructionPointer+3))
		fmt.Printf(" ip: %d\n", i.instructionPointer)
		fmt.Printf(" rb: %d\n", i.relativeBase)
		fmt.Printf("%8d[%s] -> %d\n", i.mget(i.instructionPointer+1), modes.getModeS(0), param1Value)
		fmt.Printf("%8d[%s] -> %d\n", i.mget(i.instructionPointer+2), modes.getModeS(1), param2Value)
		fmt.Printf("%8d[%s] -> m[%d] = %d\n", i.mget(i.instructionPointer+3), modes.getModeS(2), outputIndex, param1Value*param2Value)
	}

	i.mput(outputIndex, param1Value*param2Value)
	i.instructionPointer += 4
}

func (i *Intcode) halt() {
	if debug {
		fmt.Printf("\nhlt: %d\n", i.mget(i.instructionPointer))
		fmt.Printf(" ip: %d\n", i.instructionPointer)
		fmt.Printf(" rb: %d\n", i.relativeBase)
	}

	i.instructionPointer += 1
}

func (i *Intcode) input(modes *modes) {
	inputIndex := i.valueIndex(0, modes, i.instructionPointer+1)
	inputValue := i.in.Get()

	if debug {
		fmt.Printf("\ninp: %d %d\n", i.mget(i.instructionPointer), i.mget(i.instructionPointer+1))
		fmt.Printf(" ip: %d\n", i.instructionPointer)
		fmt.Printf(" rb: %d\n", i.relativeBase)
		fmt.Printf("%8d[%s] -> m[%d] = %d\n", i.mget(i.instructionPointer+1), modes.getModeS(0), inputIndex, inputValue)
	}

	i.mput(inputIndex, inputValue)
	i.instructionPointer += 2
}

func (i *Intcode) output(modes *modes) int {
	outputValue := i.value(0, modes, i.instructionPointer+1)
	i.out = outputValue

	if debug {
		fmt.Printf("\nout: %d %d\n", i.mget(i.instructionPointer), i.mget(i.instructionPointer+1))
		fmt.Printf(" ip: %d\n", i.instructionPointer)
		fmt.Printf(" rb: %d\n", i.relativeBase)
		fmt.Printf("%8d[%s] -> %d\n", i.mget(i.instructionPointer+1), modes.getModeS(0), outputValue)
	}

	i.instructionPointer += 2
	return outputValue
}

func (i *Intcode) jumpIfTrue(modes *modes) {
	param1Value := i.value(0, modes, i.instructionPointer+1)
	param2Value := i.value(1, modes, i.instructionPointer+2)

	if debug {
		fmt.Printf("\njit: %d %d %d\n", i.mget(i.instructionPointer), i.mget(i.instructionPointer+1), i.mget(i.instructionPointer+2))
		fmt.Printf(" ip: %d\n", i.instructionPointer)
		fmt.Printf(" rb: %d\n", i.relativeBase)
		fmt.Printf("%8d[%s] -> %d\n", i.mget(i.instructionPointer+1), modes.getModeS(0), param1Value)
		fmt.Printf("%8d[%s] -> %d\n", i.mget(i.instructionPointer+2), modes.getModeS(1), param2Value)
	}

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

	if debug {
		fmt.Printf("\njif: %d %d %d\n", i.mget(i.instructionPointer), i.mget(i.instructionPointer+1), i.mget(i.instructionPointer+2))
		fmt.Printf(" ip: %d\n", i.instructionPointer)
		fmt.Printf(" rb: %d\n", i.relativeBase)
		fmt.Printf("%8d[%s] -> %d\n", i.mget(i.instructionPointer+1), modes.getModeS(0), param1Value)
		fmt.Printf("%8d[%s] -> %d\n", i.mget(i.instructionPointer+2), modes.getModeS(1), param2Value)
	}

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
	outputIndex := i.valueIndex(2, modes, i.instructionPointer+3)

	result := 0
	if param1Value < param2Value {
		result = 1
	}

	if debug {
		fmt.Printf("\n lt: %d %d %d %d\n", i.mget(i.instructionPointer), i.mget(i.instructionPointer+1), i.mget(i.instructionPointer+2), i.mget(i.instructionPointer+3))
		fmt.Printf(" ip: %d\n", i.instructionPointer)
		fmt.Printf(" rb: %d\n", i.relativeBase)
		fmt.Printf("%8d[%s] -> %d\n", i.mget(i.instructionPointer+1), modes.getModeS(0), param1Value)
		fmt.Printf("%8d[%s] -> %d\n", i.mget(i.instructionPointer+2), modes.getModeS(1), param2Value)
		fmt.Printf("%8d[%s] -> m[%d] = %d\n", i.mget(i.instructionPointer+3), modes.getModeS(2), outputIndex, result)
	}

	i.mput(outputIndex, result)
	i.instructionPointer += 4
}

func (i *Intcode) equals(modes *modes) {
	// equals: if the first parameter is equal to the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
	param1Value := i.value(0, modes, i.instructionPointer+1)
	param2Value := i.value(1, modes, i.instructionPointer+2)
	outputIndex := i.valueIndex(2, modes, i.instructionPointer+3)

	result := 0
	if param1Value == param2Value {
		result = 1
	}

	if debug {
		fmt.Printf("\n eq: %d %d %d %d\n", i.mget(i.instructionPointer), i.mget(i.instructionPointer+1), i.mget(i.instructionPointer+2), i.mget(i.instructionPointer+3))
		fmt.Printf(" ip: %d\n", i.instructionPointer)
		fmt.Printf(" rb: %d\n", i.relativeBase)
		fmt.Printf("%8d[%s] -> %d\n", i.mget(i.instructionPointer+1), modes.getModeS(0), param1Value)
		fmt.Printf("%8d[%s] -> %d\n", i.mget(i.instructionPointer+2), modes.getModeS(1), param2Value)
		fmt.Printf("%8d[%s] -> %d -> %d\n", i.mget(i.instructionPointer+3), modes.getModeS(2), outputIndex, result)
	}

	i.mput(outputIndex, result)
	i.instructionPointer += 4
}

func (i *Intcode) rbOffset(modes *modes) {
	rbOffset := i.value(0, modes, i.instructionPointer+1)

	if debug {
		fmt.Printf("\nrbo: %d %d\n", i.mget(i.instructionPointer), i.mget(i.instructionPointer+1))
		fmt.Printf(" ip: %d\n", i.instructionPointer)
		fmt.Printf(" rb: %d\n", i.relativeBase)
		fmt.Printf("%8d[%s] -> %d\n", i.mget(i.instructionPointer+1), modes.getModeS(0), rbOffset)
	}

	i.relativeBase += rbOffset
	i.instructionPointer += 2
}

func (i *Intcode) value(paramIndex int, modes *modes, instructionPointer int) int {
	return i.mget(i.valueIndex(paramIndex, modes, instructionPointer))
}

func (i *Intcode) valueIndex(paramIndex int, modes *modes, instructionPointer int) int {
	switch modes.getMode(paramIndex) {
	case position:
		return i.mget(instructionPointer)
	case immediate:
		return instructionPointer
	case relative:
		return i.relativeBase + i.mget(instructionPointer)
	}
	panic(fmt.Sprintf("%d mode unknown!\n", modes.getMode(paramIndex)))
}
