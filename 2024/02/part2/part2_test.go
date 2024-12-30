package part2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckReport(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{
			input:    []int{7, 7, 4, 2, 1},
			expected: true,
		},
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
			expected: true,
		},
		{
			input:    []int{8, 6, 4, 4, 1},
			expected: true,
		},
		{
			input:    []int{1, 3, 6, 7, 9},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			output := checkReport(tt.input)
			assert.Equal(t, tt.expected, output)
		})
	}
}

func TestCheckDirection(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{
			input:    []int{7, 7, 4, 2, 1},
			expected: false,
		},
		{
			input:    []int{7, 6, 4, 2, 1},
			expected: false,
		},
		{
			input:    []int{1, 2, 7, 8, 9},
			expected: true,
		},
		{
			input:    []int{9, 7, 6, 2, 1},
			expected: false,
		},
		{
			input:    []int{1, 3, 2, 4, 5},
			expected: true,
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
			output := checkDirection(tt.input)
			assert.Equal(t, tt.expected, output)
		})
	}
}
