package part1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllOneDirection(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			input:    []int{-1, -3, -2, -4, -5},
			expected: true,
		},
		{
			input:    []int{-1, 3, -2, 4, -5},
			expected: false,
		},
		{
			input:    []int{-1, 3, -2, 4, -4},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			output := allOneDirection(tt.input)
			assert.Equal(t, tt.expected, output)
		})
	}
}

func TestGentleSlope(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{
			input:    []int{1, 2, 3, 3, 3},
			expected: true,
		},
		{
			input:    []int{-1, -3, -2, -1, -2},
			expected: true,
		},
		{
			input:    []int{1, 3, 2, 4, 5},
			expected: false,
		},
		{
			input:    []int{8, 6, 4, 4, 1},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			output := gentleSlope(tt.input, 3)
			assert.Equal(t, tt.expected, output)
		})
	}
}
