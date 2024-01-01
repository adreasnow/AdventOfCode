package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type lightPos struct {
	x int
	y int
}

type instruction struct {
	lightRange []lightPos
	action     string
}

func ReadStrings(fileName string) ([]string, error) {
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

func CountLights(lights *[1000][1000]int) int {
	count := 0
	for _, i := range lights {
		for _, j := range i {
			count += j
		}
	}
	return count
}

func BuildRange(l lightPos, u lightPos, r *instruction) {
	for x := l.x; x <= u.x; x++ {
		for y := l.y; y <= u.y; y++ {
			r.lightRange = append(r.lightRange, lightPos{x: x, y: y})
		}
	}
}

func BuildInstruction(s string) (*instruction, error) {
	inst := instruction{}
	l := len(strings.Split(s, " "))
	lowerpos := lightPos{}
	upperpos := lightPos{}
	var schema string

	if l == 5 {
		schema = "turn %s %d,%d through %d,%d"
	} else if l == 4 {
		schema = "%s %d,%d through %d,%d"
	}
	_, err := fmt.Sscanf(s, schema, &inst.action, &lowerpos.x, &lowerpos.y, &upperpos.x, &upperpos.y)
	if err != nil {
		return nil, err
	}
	BuildRange(lowerpos, upperpos, &inst)
	return &inst, nil
}

func EnactInstruction(inst *instruction, lights *[1000][1000]int) {
	for _, i := range inst.lightRange {
		if inst.action == "on" {
			lights[i.x][i.y]++
		} else if inst.action == "off" {
			if lights[i.x][i.y] != 0 {
				lights[i.x][i.y]--
			}
		} else if inst.action == "toggle" {
			lights[i.x][i.y] += 2
		}
	}
}

func main() {
	lights := [1000][1000]int{}
	stringList, err := ReadStrings("input.txt")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, s := range stringList {
			inst, err := BuildInstruction(s)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				EnactInstruction(inst, &lights)
			}
		}
		fmt.Println(CountLights(&lights))
	}
}
