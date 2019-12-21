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
#ABCD#

   ABCD
  J123
####.###

  ABCD
 J123
##...###

     ABCD
 ABCD
J123J123
###.##.###


*/

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)
	cpu := intcode.NewIntcode(intcode.NewMemory(lines[0]))
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
		fmt.Println(v)
		fmt.Printf("%c", r)
	}
}
