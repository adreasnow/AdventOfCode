package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInts(fileName string) []int {
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("could not open file", err)
		os.Exit(1)
	}

	input := make([]int, 0)

	for _, line := range strings.Split(string(file), "\n") {
		if line != "" {
			val, err := strconv.Atoi(string(line))
			if err != nil {
				fmt.Println("could not parse line:", err)
				os.Exit(1)
			}
			input = append(input, val)
		}
	}

	return input
}
