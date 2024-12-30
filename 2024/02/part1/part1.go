package part1

import "fmt"

func calculateSteps(report []int) []int {
	steps := []int{}
	for n := range report {
		if n != len(report)-1 {
			step := report[n+1] - report[n]
			steps = append(steps, step)
		}
	}
	return steps
}

func determineSafety(report []int) bool {
	steps := calculateSteps(report)

	switch {
	case !allOneDirection(steps):
		return false
	case !gentleSlope(steps, 3):
		return false
	}

	return true
}

func Part1(reports [][]int) {
	safeCount := 0
	for _, report := range reports {
		if determineSafety(report) {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}
