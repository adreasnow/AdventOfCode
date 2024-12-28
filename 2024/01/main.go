package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	fmt.Println("Part 1:")
	part1()
	fmt.Println("")
	fmt.Println("Part 2:")
	part2()

}

func part1() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("could not read input file: %s", err.Error())
		os.Exit(1)
	}

	list1 := []int{}
	list2 := []int{}

	for _, s := range strings.Split(string(bytes), "\n") {
		var n1 int
		var n2 int

		_, err := fmt.Sscanf(s, "%d   %d", &n1, &n2)
		if err == nil {
			list1 = append(list1, n1)
			list2 = append(list2, n2)
		}
	}

	slices.Sort(list1)
	slices.Sort(list2)

	distance := 0
	for i := range list1 {
		var d int
		d = list1[i] - list2[i]
		if d < 0 {
			d = list2[i] - list1[i]
		}
		distance += d
	}
	fmt.Println(distance)
}

func part2() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("could not read input file: %s", err.Error())
		os.Exit(1)
	}

	list1 := []int{}
	list2 := []int{}

	for _, s := range strings.Split(string(bytes), "\n") {
		var n1 int
		var n2 int

		_, err := fmt.Sscanf(s, "%d   %d", &n1, &n2)
		if err == nil {
			list1 = append(list1, n1)
			list2 = append(list2, n2)
		}
	}

	countMap := map[int]int{}

	for _, n := range list2 {
		countMap[n]++
	}

	similarity := 0
	for _, n := range list1 {
		similarity += n * countMap[n]
	}

	fmt.Println(similarity)

}
