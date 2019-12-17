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
	fmt.Printf("Read %d lines!\n", len(lines))
	cpu := intcode.NewIntcode(intcode.NewMemory(lines[0]))
	grid := [][]string{}
	row := []string{}
	for {
		o, m := cpu.Output()

		if m == intcode.HaltMode {
			break
		}
		switch o {
		case '#':
			fmt.Print("#")
			row = append(row, "#")
		case '.':
			fmt.Print(".")
			row = append(row, ".")
		case '\n':
			fmt.Println()
			grid = append(grid, row)
			row = []string{}
		case '>':
			fmt.Print(">")
			row = append(row, ">")
		case '<':
			fmt.Print("<")
			row = append(row, "<")
		case '^':
			fmt.Print("^")
			row = append(row, "^")
		case 'v':
			fmt.Print("v")
			row = append(row, "v")
		default:
			fmt.Print("?")
			row = append(row, ">")
		}
	}

	sum := 0
	for y, row := range grid {
		for x, tile := range row {
			if x == 0 || y == 0 || y == len(grid)-1 || x == len(row)-1 {
				fmt.Print(tile)
				continue
			}
			if tile == "#" && grid[y][x-1] == "#" && grid[y][x+1] == "#" && grid[y-1][x] == "#" && grid[y+1][x] == "#" {
				sum += x * y
				fmt.Print("O")
			} else {
				fmt.Print(tile)
			}
		}
		fmt.Println()
	}

	// 4111 too highkkkj
	fmt.Println(sum)
}
