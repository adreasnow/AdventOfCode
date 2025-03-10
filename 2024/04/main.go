package main

import (
	"fmt"
	"os"
	"strings"
)

func parseInput(input string) *data {
	d := make(data, 0)
	for _, line := range strings.Split(input, "\n") {
		lr := make([]rune, len(line))
		for j, r := range line {
			lr[j] = r
		}
		if len(lr) > 1 {
			d = append(d, lr)
		}
	}

	return &d
}

func main() {
	input := readInput("input.txt")
	d := parseInput(input)

	fmt.Println("Part 1:")
	count := part1(d)
	fmt.Println(count)
	fmt.Println("-----------")

	fmt.Println("Part 2:")
	count = part2(d)
	fmt.Println(count)
	fmt.Println("Done")
}

func part1(d *data) int {
	return d.iterateIndexes(1)
}

func part2(d *data) int {
	return d.iterateIndexes(2)
}

func readInput(file string) string {
	body, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("could not read input file")
	}
	return string(body)
}
