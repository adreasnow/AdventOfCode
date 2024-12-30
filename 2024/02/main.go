package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"aoc/2024/02/part1"
	"aoc/2024/02/part2"
)

func loadReports() [][]int {
	reports := [][]int{}

	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("could not read input file: %s", err.Error())
		os.Exit(1)
	}

	for _, line := range strings.Split(string(bytes), "\n") {
		report := []int{}
		if line != "" {
			for _, level := range strings.Split(line, " ") {
				levelInt, err := strconv.Atoi(level)
				if err != nil {
					fmt.Printf("error converting string to int: %v\n", err)
					os.Exit(1)
				}

				report = append(report, levelInt)
			}
			reports = append(reports, report)
		}
	}
	return reports
}

func main() {
	reports := loadReports()
	fmt.Println("Part 1:")
	part1.Part1(reports)
	fmt.Println("---------")
	fmt.Println("Part 2:")
	part2.Part2(reports)
	fmt.Println("---------")
}
