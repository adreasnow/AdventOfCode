package main

import (
	"fmt"
	"os"
)

func calcFloor(input string) (int, error) {
	floor := 0
	o := byte('(')
	c := byte(')')

	data, err := os.ReadFile(input)
	if err != nil {
		return 0, err
	}

	for _, b := range data {
		if b == o {
			floor++
		} else if b == c {
			floor--
		}
	}
	return floor, nil
}

func main() {
	floor, err := calcFloor("input.txt")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(floor)
	}
}
