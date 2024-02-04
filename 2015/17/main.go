package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadStrings(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	input := make([]int, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return input, err
		}
		input = append(input, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return input, err
}

// Lets have some fun and writre a binary adder manually :)
func BinCountUp(b []bool) []bool {
	carry := false
	add := make([]bool, len(b))
	add[0] = true
	for i := 0; i < len(b); i++ {
		if b[i] && add[i] {
			b[i] = false
			carry = true
		} else if (!b[i] && add[i]) || (!add[i] && b[i]) {
			if !carry {
				b[i] = true
				carry = false
			} else {
				b[i] = false
				carry = true
			}
		} else if !b[i] && !add[i] {
			if !carry {
				b[i] = false
				carry = false
			} else {
				b[i] = true
				carry = false
			}
		}
	}
	return b
}

func SumBoolSlice(s []bool) int {
	total := 0
	for _, check := range s {
		if check {
			total++
		}
	}
	return total
}

func GeneratePermutations(n int) [][]bool {
	permutations := make([][]bool, 0)
	perm := make([]bool, n)
	for SumBoolSlice(perm) != len(perm) {
		perm := BinCountUp(perm)
		permutations = append(permutations, []bool{})
		// Go really doesn't like this, but I don't know how else to get it to append the values and not the reference
		for _, i := range perm {
			permutations[len(permutations)-1] = append(permutations[len(permutations)-1], i)
		}
	}
	return permutations
}

func CountCombinations(perms [][]bool, buckets []int) int {
	var total int
	minContainers := len(buckets)

	count := 0
	minCount := 0

	for _, perm := range perms {
		total = 0
		for i, check := range perm {
			if check {
				total += buckets[i]
			}
		}
		if total == 150 {
			count++
			if SumBoolSlice(perm) < minContainers {
				minContainers = SumBoolSlice(perm)
			}
		}
	}
	for _, perm := range perms {
		total = 0
		for i, check := range perm {
			if check {
				total += buckets[i]
			}
		}
		if total == 150 && SumBoolSlice(perm) == minContainers {
			minCount++
		}
	}
	return minCount
}

func main() {

	buckets, err := ReadStrings("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	perms := GeneratePermutations(len(buckets))
	combinations := CountCombinations(perms, buckets)
	fmt.Printf("There are %d combinations of containers that add up to 150L with the smallest number of containers", combinations)

}
