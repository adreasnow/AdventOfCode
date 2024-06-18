package main

import (
	"bufio"
	"fmt"
	"os"
)

func readStrings(fileName string) ([]string, error) {
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

func (i *Instructions) parseInput() {
	lines, err := readStrings("input.txt")
	if err != nil {
		fmt.Println("Could not read from file", err)
		os.Exit(1)
	}

	for _, line := range lines {
		var r string
		var o string

		if line == "" {
			return
		}

		switch inst := line[0:3]; inst {
		case "hlf":
			fmt.Sscanf(line, "hlf %s", &r)
			(*i).add(hlf(r))

		case "tpl":
			fmt.Sscanf(line, "tpl %s", &r)
			(*i).add(tpl(r))

		case "inc":
			fmt.Sscanf(line, "inc %s", &r)
			(*i).add(inc(r))

		case "jmp":
			fmt.Sscanf(line, "jmp %s", &o)
			(*i).add(jmp(o))

		case "jie":
			fmt.Sscanf(line, "jie %s %s", &r, &o)
			r = string(r[0])
			(*i).add(jie(r, o))

		case "jio":
			fmt.Sscanf(line, "jio %s %s", &r, &o)
			r = string(r[0])
			(*i).add(jio(r, o))
		}
	}
	return
}
