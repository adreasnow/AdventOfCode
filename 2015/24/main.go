package main

import (
	"fmt"
	"slices"

	"github.com/ernestosuarez/itertools"
)

// 1. three equally weighted groups
// 2. group 1 - minimum ampout of packages
// 3. if multiple options - first has smalles QE

func sumSlice(s []int) int {
	sum := 0
	for _, pkg := range s {
		sum += pkg
	}
	return sum
}

func solve(packageGroups int) int {
	packages := readInts("input.txt")
	slices.Reverse(packages)

	packagesWeight := sumSlice(packages)
	targetSize := packagesWeight / packageGroups

	comboList := [][]int{}
	qe := -1
	packageCount := 1
	for qe < 0 {
		for combination := range itertools.CombinationsInt(packages, packageCount) {
			if sumSlice(combination) == targetSize {
				comboList = append(comboList, combination)
			}
		}
		qe = smallestQE(comboList)
		packageCount++
	}
	return qe
}

func main() {
	fmt.Println("Part 1:")
	qe := solve(3)
	fmt.Println("-----------")
	fmt.Println(qe)
	fmt.Println("")
	fmt.Println("Part 2:")
	qe = solve(4)
	fmt.Println("-----------")
	fmt.Println(qe)

}
