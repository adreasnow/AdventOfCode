package main

import (
	"fmt"
	"reflect"
	"testing"
)

// func TestCalculatePresentsPartA(t *testing.T) {
// 	presents := make([]int, 0)
// 	for i := range 9 {
// 		presents = append(presents, CalculatePresents(i+1))
// 	}

// 	expected := []int{10, 30, 40, 70, 60, 120, 80, 150, 130}

// 	if !reflect.DeepEqual(presents, expected) {
// 		fmt.Println("Output: ", presents)
// 		fmt.Println("Expected: ", expected)
// 		t.Errorf("The output did not match expectation")
// 	}
// }

func TestCalculatePresentsPartB(t *testing.T) {
	presents := make([]int, 0)
	elves := make(map[int]int, 0)

	for i := range 9 {
		presents = append(presents, CalculatePresents(i+1, &elves))
	}

	expected := []int{11, 33, 44, 77, 66, 132, 88, 165, 143}

	if !reflect.DeepEqual(presents, expected) {
		fmt.Println("Output: ", presents)
		fmt.Println("Expected: ", expected)
		t.Errorf("The output did not match expectation")
	}
}
