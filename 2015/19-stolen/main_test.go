package main

import (
	"fmt"
	"testing"
)

func TestRemoveIndex(t *testing.T) {
	a := []string{"a", "b"}
	b := []string{"c", "d"}
	input := [][]string{a, b}

	output := RemoveIndex(input, b)
	expected := [][]string{a}

	if len(output) != 1 {
		t.Errorf("Output of did not match expectation.")
		fmt.Println("Output", output)
		fmt.Println("Expected", expected)
	}
}

func TestGetLongestIndex(t *testing.T) {
	a := []string{"a", "bbb"}
	b := []string{"c", "dd"}
	c := []string{"e", "ffff"}
	input := [][]string{a, b, c}

	output, _ := GetLongestIndex(input)
	expected := c

	if output[0] != c[0] {
		t.Errorf("Output of did not match expectation.")
		fmt.Println("Output", output)
		fmt.Println("Expected", expected)
	}
}
