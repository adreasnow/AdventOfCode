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

func buildMap(data []byte) map[string]int {
	coordinates := make(map[string]int)

	x1 := 0
	y1 := 0
	x2 := 0
	y2 := 0
	u := byte('^')
	d := byte('v')
	r := byte('>')
	l := byte('<')

	var key string

	for count, b := range data {
		if count%2 == 0 {
			if b == u {
				y1++
			} else if b == d {
				y1--
			} else if b == l {
				x1--
			} else if b == r {
				x1++
			}
			key = fmt.Sprintf("%05d%05d", x1, y1)
			coordinates[key] = 1
		} else {
			if b == u {
				y2++
			} else if b == d {
				y2--
			} else if b == l {
				x2--
			} else if b == r {
				x2++
			}
			key = fmt.Sprintf("%05d%05d", x2, y2)
			coordinates[key] = 1
		}
	}
	return coordinates
}

func main() {
	data, err := readData("input.txt")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		coordinates := buildMap(data)
		fmt.Printf("%d houses have been visited.\n", len(coordinates))
	}
}
