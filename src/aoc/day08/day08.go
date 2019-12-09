package main

import (
	"aoc"
	"fmt"
)

func main() {
	inputFilePath := aoc.InputArg()
	Main(inputFilePath)
}

func Main(inputFilePath string) {
	lines := aoc.Read(inputFilePath)
	layers := toLayers(lines[0])
	fmt.Println("Exercise 1:", exercise1(layers))

	fmt.Println("Exercise 2:")
	merged := exercise2(layers)
	for i := 0; i < layerPixels; i += lineWidth {
		fmt.Println(string(merged[i : i+lineWidth]))
	}
}

const (
	black       = '0'
	white       = '1'
	transparent = '2'
	lineWidth   = 25
	layerPixels = lineWidth * 6
)

type layer []rune

func toLayers(line string) []layer {
	layers := []layer{}
	for start := 0; start < len(line); start += layerPixels {
		layerstr := line[start : start+layerPixels]
		layers = append(layers, []rune(layerstr))
	}
	return layers
}

func exercise1(layers []layer) int {
	value := 0
	lowestZerosCount := layerPixels

	for _, l := range layers {
		zeros := 0
		ones := 0
		twos := 0

		for _, r := range l {
			switch r {
			case '0':
				zeros++
			case '1':
				ones++
			case '2':
				twos++
			}
		}

		if zeros < lowestZerosCount {
			value = ones * twos
			lowestZerosCount = zeros
		}
	}

	return value
}

func exercise2(layers []layer) layer {
	merged := layers[0]
	for _, l := range layers[1:] {
		for i, c := range l {
			if merged[i] == transparent {
				merged[i] = c
			}
		}
	}
	for i, c := range merged {
		switch c {
		case white:
			merged[i] = '#'
		case black:
			merged[i] = ' '
		}
	}
	return merged
}
