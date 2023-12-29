package main

import (
	"fmt"
	"os"
)

func readData(input string) ([]byte, error) {
	data, err := os.ReadFile(input)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func calcFloor(data []byte) (int, int) {
	floor := 0
	basement := 0
	o := byte('(')
	c := byte(')')

	for count, b := range data {
		if b == o {
			floor++
		} else if b == c {
			floor--
		}
		if floor == -1 && basement == 0 {
			basement = count + 1
		}
	}
	return floor, basement
}

func main() {
	data, err := readData("input.txt")
	floor, basement := calcFloor(data)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("The final floor is %d\n", floor)
		fmt.Printf("Character that causes Santa to enter the basement is: %d", basement)
	}
}
