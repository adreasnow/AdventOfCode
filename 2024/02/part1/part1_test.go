package part1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateSteps(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 1, 1, 1},
		},
		{
			input:    []int{1, 4, 6, 9, 12},
			expected: []int{3, 2, 3, 3},
		},
		{
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{-1, -1, -1, -1},
		},
		{
			input:    []int{10, 5, 3, 1},
			expected: []int{-5, -2, -2},
		},
		{
			input:    []int{10, 20, 5, 1},
			expected: []int{10, -15, -4},
		},
		{
			input:    []int{0, 0, 0, 0, 0},
			expected: []int{0, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			output := calculateSteps(tt.input)
			assert.Equal(t, tt.expected, output)
		})
	}
}

func TestDetermineSafety(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{
			input:    []int{7, 6, 4, 2, 1},
			expected: true,
		},
		{
			input:    []int{1, 2, 7, 8, 9},
			expected: false,
		},
		{
			input:    []int{9, 7, 6, 2, 1},
			expected: false,
		},
		{
			input:    []int{1, 3, 2, 4, 5},
			expected: false,
		},
		{
			input:    []int{8, 6, 4, 4, 1},
			expected: false,
		},
		{
			input:    []int{1, 3, 6, 7, 9},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			output := determineSafety(tt.input)
			assert.Equal(t, tt.expected, output)
		})
	}
}
