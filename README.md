# Advent of Code 2019

My best attempts to solve [Advent of Code 2019](https://adventofcode.com/2019) in Go.

## Development

Scaffold code for next day could be created with `./scallfold.sh`. Based on `src/aoc` structure next day would be created.

When code is ready it could be executed with:

    GOPATH=`pwd` go run src/day00/day00.go input/day00_input.txt

Each day could be treated as a plugin. To use it this way one would need to build plugins:

    ./build_plugins.sh

This command will create plugins from each day and place them in `plugins` directory. Then each of the day could be executed with shared `main`:

    GOPATH=`pwd` go run src/main.go day00 input/day00_input.txt

## Tests

One could execute unit tests with:

    GOPATH=`pwd` go test aoc/day00
