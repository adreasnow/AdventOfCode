package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readInstructions(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	input := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return input, err
}

type Action uint8

const (
	ASSIGN Action = iota
	AND
	OR
	LSHIFT
	RSHIFT
	NOT
)

var (
	wires = make(map[string]Wire)
	cache = make(map[string]uint16)
)

type Wire struct {
	action Action
	src1   string
	src2   string
}

func compute(target string) uint16 {
	if value, err := strconv.ParseUint(target, 10, 16); err == nil {
		return uint16(value)
	}
	if value, ok := cache[target]; ok {
		return value
	}

	wire := wires[target]
	var value uint16

	switch wire.action {
	case ASSIGN:
		value = compute(wire.src1)
	case AND:
		value = compute(wire.src1) & compute(wire.src2)
	case OR:
		value = compute(wire.src1) | compute(wire.src2)
	case LSHIFT:
		value = compute(wire.src1) << compute(wire.src2)
	case RSHIFT:
		value = compute(wire.src1) >> compute(wire.src2)
	case NOT:
		value = ^compute(wire.src1)
	}
	cache[target] = value
	return value
}

func parseInstruction(s string) error {
	var target string
	i := Wire{}

	if n, _ := fmt.Sscanf(s, "%s -> %s", &i.src1, &target); n == 2 {
		i.action = ASSIGN
	} else if n, _ := fmt.Sscanf(s, "NOT %s -> %s", &i.src1, &target); n == 2 {
		i.action = NOT
	} else if n, _ := fmt.Sscanf(s, "%s AND %s -> %s", &i.src1, &i.src2, &target); n == 3 {
		i.action = AND
	} else if n, _ := fmt.Sscanf(s, "%s OR %s -> %s", &i.src1, &i.src2, &target); n == 3 {
		i.action = OR
	} else if n, _ := fmt.Sscanf(s, "%s LSHIFT %s -> %s", &i.src1, &i.src2, &target); n == 3 {
		i.action = LSHIFT
	} else if n, _ := fmt.Sscanf(s, "%s RSHIFT %s -> %s", &i.src1, &i.src2, &target); n == 3 {
		i.action = RSHIFT
	}

	wires[target] = i
	return nil
}

func main() {
	instructions, err := readInstructions("input.txt")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, i := range instructions {
			err := parseInstruction(i)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}
	}

	parseInstruction("3176 -> b")
	// for i, res := range wires {
	// 	fmt.Printf("%s: %v\n", i, res)
	// }

	fmt.Printf("%d\n", compute("a"))
}
