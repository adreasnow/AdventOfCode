package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() (int, int) {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("could not read input file: %v", err.Error())
		os.Exit(1)
	}

	split := strings.Split(string(bytes), " ")
	row := split[16][0:4]
	col := split[18][0:4]

	rowInt, err := strconv.Atoi(row)
	if err != nil {
		fmt.Printf("could not parse row int: %v", err.Error())
		os.Exit(1)
	}
	colInt, err := strconv.Atoi(col)
	if err != nil {
		fmt.Printf("could not parse col int: %v", err.Error())
		os.Exit(1)
	}
	return rowInt, colInt
}

func solver() int {
	row, col := readInput()
	return getValueOfIndex(row, col)
}

func main() {
	fmt.Println("Part 1:")
	fmt.Println(solver())
	fmt.Println("-----------")
}
