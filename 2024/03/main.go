package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := readInput()
	fmt.Println("Part 1:")
	part1(input)
	fmt.Println("-----------")
	fmt.Println("Part 2:")
	part2(input)
	fmt.Println("Done")

}

func part1(input string) {
	sum := 0
	for _, s := range strings.Split(input, "mul(") {
		var val1, val2 int
		if n, err := fmt.Sscanf(s, "%d,%d)", &val1, &val2); n == 2 && err == nil {
			fmt.Println(val1, val2, s)
			sum += val1 * val2

		}
	}

	fmt.Println(sum)
}

func part2(input string) {
	enabled := true
	sum := 0

	for n := range input {
		s := input[n:]
		if len(s) > 6 {
			switch {
			case s[0:4] == "do()":
				enabled = true
				fmt.Println(s[0:4])
			case s[0:7] == "don't()":
				enabled = false
				fmt.Println(s[0:7])
			case s[0:4] == "mul(":
				var val1, val2 int
				if n, err := fmt.Sscanf(s, "mul(%d,%d)", &val1, &val2); n == 2 && err == nil {
					if enabled {
						sum += val1 * val2
					}
				}
			}
		}
	}
	fmt.Println(sum)
}

func readInput() string {
	body, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("could not read input file")
	}
	return string(body)
}
