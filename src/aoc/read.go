package aoc

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func InputArg() string {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <input path>\n", os.Args[0])
		os.Exit(1)
	}
	return os.Args[1]
}

func Read(filePath string) []string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
