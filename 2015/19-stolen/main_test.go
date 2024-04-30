package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveIndex(t *testing.T) {
	a := []string{"a", "b"}
	b := []string{"c", "d"}
	c := []string{"c", "e"}
	input := [][]string{a, b, c}

	output := removeIndex(input, b)
	expected := [][]string{a, c}

	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Output of did not match expectation.")
		fmt.Println("Output", output)
		fmt.Println("Expected", expected)
	}
}

func TestGetLongestIndex(t *testing.T) {
	a := []string{"a", "bbb"}
	b := []string{"c", "dd"}
	c := []string{"e", "ffff"}
	d := []string{"e", "fffg"}
	input := [][]string{a, b, c, d}

	output, _ := getLongestIndex(input)
	expected := c

	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Output of did not match expectation.")
		fmt.Println("Output", output)
		fmt.Println("Expected", expected)
	}
}
