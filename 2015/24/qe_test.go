package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalulateQE(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{11, 9},
			expected: 99,
		},
		{
			input:    []int{10, 9, 1},
			expected: 90,
		},
		{
			input:    []int{10, 8, 2},
			expected: 160,
		},
		{
			input:    []int{10, 7, 3},
			expected: 210,
		},
		{
			input:    []int{10, 5, 4, 1},
			expected: 200,
		},
		{
			input:    []int{10, 5, 3, 2},
			expected: 300,
		},
		{
			input:    []int{10, 4, 3, 2, 1},
			expected: 240,
		},
		{
			input:    []int{9, 8, 3},
			expected: 216,
		},
		{
			input:    []int{9, 7, 4},
			expected: 252,
		},
		{
			input:    []int{9, 5, 4, 2},
			expected: 360,
		},
		{
			input:    []int{8, 7, 5},
			expected: 280,
		},
		{
			input:    []int{8, 5, 4, 3},
			expected: 480,
		},
		{
			input:    []int{7, 5, 4, 3, 1},
			expected: 420,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			output := calculateQE(tt.input)
			assert.Equal(t, tt.expected, output)
		})
	}
}

func TestSmallestQE(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected int
	}{
		{
			input: [][]int{
				{11, 9},
				{10, 9, 1},
				{10, 8, 2},
				{10, 7, 3},
				{10, 5, 4, 1},
				{10, 5, 3, 2},
				{10, 4, 3, 2, 1},
				{9, 8, 3},
				{9, 7, 4},
				{9, 5, 4, 2},
				{8, 7, 5},
				{8, 5, 4, 3},
				{7, 5, 4, 3, 1},
			},
			expected: 90,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			output := smallestQE(tt.input)
			assert.Equal(t, tt.expected, output)
		})
	}
}
