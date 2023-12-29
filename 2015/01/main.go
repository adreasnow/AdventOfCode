package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	input := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return input[0], err
}

func calcFloors(FloorString *string) int {
	floor := 0
	for _, rune := range *FloorString {
		if rune == '(' {
			floor++
		} else if rune == ')' {
			floor--
		}
	}

	return floor
}

func main() {
	contents, err := readInput("input.txt")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println((calcFloors(&contents)))
	}
}
