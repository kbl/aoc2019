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
	fmt.Printf("Read %d lines!\n", len(lines))
	exercise1(lines[0])
}

const (
	black       = '0'
	white       = '1'
	transparent = '2'
)

func exercise1(line string) int {
	layers := map[int][]rune{}
	howMany := 25 * 6
	sz := howMany
	i := 0
	value := 0

	for start := 0; start < len(line); start += howMany {

		layerstr := line[start : start+howMany]
		layer := []rune(layerstr)
		layers[i] = layer
		hm := 0
		ones := 0
		twos := 0

		for _, r := range layer {
			switch r {
			case '0':
				hm++
			case '1':
				ones++
			case '2':
				twos++
			}
		}

		if hm < sz {
			value = ones * twos
		}

		layers[i] = layer
		i++
	}

	fmt.Println(value)

	merged := []rune("222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222")

	for i := 0; i < len(layers); i++ {
		l := layers[i]
		for i, c := range l {
			if merged[i] == transparent {
				merged[i] = c
			}
		}
	}

	hm := 0
	ones := 0
	twos := 0

	for _, r := range merged {
		switch r {
		case '0':
			hm++
		case '1':
			ones++
		case '2':
			twos++
		}
	}

	fmt.Println(ones)

	fmt.Println(string(merged[:25]))
	fmt.Println(string(merged[25:50]))
	fmt.Println(string(merged[50:75]))
	fmt.Println(string(merged[75:100]))
	fmt.Println(string(merged[100:125]))
	fmt.Println(string(merged[125:150]))

	return 0
}
