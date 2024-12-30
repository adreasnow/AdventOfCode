package part2

import (
	"fmt"
	"math"
)

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

func checkDirection(report []int) (positive bool) {
	directions := 0

	for n := range report {
		if n == 0 {
			continue
		}

		switch {
		case report[n]-report[n-1] < 0:
			directions--
		case report[n]-report[n-1] > 0:
			directions++
		}
	}

	positive = directions >= 0
	return
}

func checkReport(report []int) bool {
	positive := checkDirection(report)
	removed := false

	for len(report) > 1 {
		fmt.Println(report)
		switch {
		// if the step is zero
		case report[1]-report[0] == 0 && removed:
			fmt.Println("zero step, fail")
			return false
		case report[1]-report[0] == 0 && !removed:
			removed = true
			fmt.Println("zero step, removing")
			report = report[1:]
			continue

		// if the step is positive and the direction is negative
		case report[1]-report[0] > 0 && !positive && removed:
			fmt.Println("wrong direction step, fail")
			return false
		case report[1]-report[0] > 0 && !positive && !removed:
			removed = true
			fmt.Println("wrong direction step, removing")

			if len(report) >= 3 {
				if report[2]-report[0] > 0 {
					report = append([]int{report[0]}, report[2:]...)
				} else {
					report = report[1:]
				}
			} else {
				report = []int{report[0]}
			}
			continue

		// if the step is negative and the direction is positive
		case report[1]-report[0] < 0 && positive && removed:
			fmt.Println("wrong direction step, fail")
			return false
		case report[1]-report[0] < 0 && positive && !removed:
			removed = true
			fmt.Println("wrong direction step, removing")

			if len(report) >= 3 {
				if report[2]-report[0] < 0 {
					report = append([]int{report[0]}, report[2:]...)
				} else {
					report = report[1:]
				}
			} else {
				report = []int{report[0]}
			}
			continue

		// if the step size is too big
		case int(math.Abs(float64(report[1]-report[0]))) > 3 && removed:
			fmt.Println("too big step, fail")
			return false
		case int(math.Abs(float64(report[1]-report[0]))) > 3 && !removed:
			fmt.Println("too big step, removing")
			removed = true

			if len(report) >= 3 {
				if int(math.Abs(float64(report[2]-report[0]))) <= 3 {
					report = append([]int{report[0]}, report[2:]...)
				} else {
					report = report[1:]
				}
			} else {
				report = []int{report[0]}
			}
			continue
		}
		report = report[1:]
	}

	return true
}

func Part2(reports [][]int) {
	total := 0

	for _, report := range reports {
		if checkReport(report) {
			total++
		}
	}

	fmt.Println(total)
}
