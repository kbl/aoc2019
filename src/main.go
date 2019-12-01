package main

import (
	"fmt"
	"log"
	"os"
	"plugin"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: %s <day00> <input path>\n", os.Args[0])
		os.Exit(1)
	}
	day := os.Args[1]
	inputPath := os.Args[2]

	dayPlugin, err := plugin.Open(fmt.Sprintf("plugins/%s.so", day))
	if err != nil {
		log.Fatal(err)
	}

	dayMain, err := dayPlugin.Lookup("Main")
	if err != nil {
		log.Fatal(err)
	}

	dayMain.(func(string))(inputPath)
}
