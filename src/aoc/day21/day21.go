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

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)
	fmt.Println("Exercise 1:", exercise1(lines[0]))
	fmt.Println("Exercise 2:", exercise2(lines[0]))
}

func exercise1(memory string) int {
	cpu := intcode.NewIntcode(intcode.NewMemory(memory))
	program := `NOT C T
OR T J
NOT A T
OR T J
AND D J
WALK
`
	for _, x := range program {
		cpu.AddInput(int(x))
	}

	output := intcode.OutputMode
	for output == intcode.OutputMode {
		v, o := cpu.Output()
		output = o
		r := rune(v)
		if v > 127 {
			return v
		}
		fmt.Printf("%c", r)
	}
	return -1
}

/*

AND X Y sets Y to true if both X and Y are true;
        otherwise, it sets Y to false.

OR X Y  sets Y to true if at least one of X or Y is true;
        otherwise, it sets Y to false.

NOT X Y sets Y to true if X is false;
        otherwise, it sets Y to false.

hole   – false
ground – true


@
#ABCDEFGHI#

   2
  1 3
 J   L
####.###
   2
  1 3
 J   L
##...###

   2    2
  1 3  1 3
 J   LJ   L
####.##.###

   2   2   2
  1 3 1 3 1 3
 J   x   x   L
####.##.##..####

##.##..####

#####.#.##..#####

#####..###...####
   ABCDEFGH

*/
func exercise2(memory string) int {
	cpu := intcode.NewIntcode(intcode.NewMemory(memory))
	program := `NOT T T
AND A T
AND B T
AND C T
NOT T J
AND D J
AND H J
NOT A T
OR T J
RUN
`
	for _, x := range program {
		cpu.AddInput(int(x))
	}

	output := intcode.OutputMode
	for output == intcode.OutputMode {
		v, o := cpu.Output()
		output = o
		r := rune(v)
		if v > 127 {
			return v
		}
		fmt.Printf("%c", r)
	}
	return -1
}
