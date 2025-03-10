package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var dummyData = data{
	[]rune{'A', 'A', 'A', 'A', 'A'},
	[]rune{'A', 'X', 'M', 'X', 'S'},
	[]rune{'A', 'M', 'M', 'M', 'A'},
	[]rune{'A', 'A', 'A', 'A', 'A'},
	[]rune{'A', 'S', 'A', 'S', 'S'},
}

func TestUp(t *testing.T) {
	tests := []struct {
		input    index
		expected [4]index
		check    bool
	}{
		{
			input: index{3, 2},
			expected: [4]index{
				{3, 2},
				{2, 2},
				{1, 2},
				{0, 2},
			},
			check: true,
		},
		{
			input:    index{2, 2},
			expected: [4]index{},
			check:    false,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			output, check := dummyData.up(tt.input)
			assert.Equal(t, tt.expected, output)
			assert.Equal(t, tt.check, check)
		})
	}
}

func TestDown(t *testing.T) {
	tests := []struct {
		input    index
		expected [4]index
		check    bool
	}{
		{
			input: index{1, 2},
			expected: [4]index{
				{1, 2},
				{2, 2},
				{3, 2},
				{4, 2},
			},
			check: true,
		},
		{
			input:    index{2, 2},
			expected: [4]index{},
			check:    false,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			output, check := dummyData.down(tt.input)
			assert.Equal(t, tt.expected, output)
			assert.Equal(t, tt.check, check)
		})
	}
}

func TestRight(t *testing.T) {
	tests := []struct {
		input    index
		expected [4]index
		check    bool
	}{
		{
			input: index{1, 1},
			expected: [4]index{
				{1, 1},
				{1, 2},
				{1, 3},
				{1, 4},
			},
			check: true,
		},
		{
			input:    index{1, 2},
			expected: [4]index{},
			check:    false,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			output, check := dummyData.right(tt.input)
			assert.Equal(t, tt.expected, output)
			assert.Equal(t, tt.check, check)
		})
	}
}

func TestLeft(t *testing.T) {
	tests := []struct {
		input    index
		expected [4]index
		check    bool
	}{
		{
			input: index{1, 3},
			expected: [4]index{
				{1, 3},
				{1, 2},
				{1, 1},
				{1, 0},
			},
			check: true,
		},
		{
			input:    index{1, 2},
			expected: [4]index{},
			check:    false,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			output, check := dummyData.left(tt.input)
			assert.Equal(t, tt.expected, output)
			assert.Equal(t, tt.check, check)
		})
	}
}

func TestUpLeft(t *testing.T) {
	tests := []struct {
		input    index
		expected [4]index
		check    bool
	}{
		{
			input: index{3, 3},
			expected: [4]index{
				{3, 3},
				{2, 2},
				{1, 1},
				{0, 0},
			},
			check: true,
		},
		{
			input:    index{3, 2},
			expected: [4]index{},
			check:    false,
		},
		{
			input:    index{2, 3},
			expected: [4]index{},
			check:    false,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			output, check := dummyData.upLeft(tt.input)
			assert.Equal(t, tt.expected, output)
			assert.Equal(t, tt.check, check)
		})
	}
}

func TestUpRight(t *testing.T) {
	tests := []struct {
		input    index
		expected [4]index
		check    bool
	}{
		{
			input: index{3, 1},
			expected: [4]index{
				{3, 1},
				{2, 2},
				{1, 3},
				{0, 4},
			},
			check: true,
		},
		{
			input:    index{2, 1},
			expected: [4]index{},
			check:    false,
		},
		{
			input:    index{3, 2},
			expected: [4]index{},
			check:    false,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			output, check := dummyData.upRight(tt.input)
			assert.Equal(t, tt.expected, output)
			assert.Equal(t, tt.check, check)
		})
	}
}

func TestDownLeft(t *testing.T) {
	tests := []struct {
		input    index
		expected [4]index
		check    bool
	}{
		{
			input: index{1, 3},
			expected: [4]index{
				{1, 3},
				{2, 2},
				{3, 1},
				{4, 0},
			},
			check: true,
		},
		{
			input:    index{2, 3},
			expected: [4]index{},
			check:    false,
		},
		{
			input:    index{1, 2},
			expected: [4]index{},
			check:    false,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			output, check := dummyData.downLeft(tt.input)
			assert.Equal(t, tt.expected, output)
			assert.Equal(t, tt.check, check)
		})
	}
}

func TestDownRight(t *testing.T) {
	tests := []struct {
		input    index
		expected [4]index
		check    bool
	}{
		{
			input: index{1, 1},
			expected: [4]index{
				{1, 1},
				{2, 2},
				{3, 3},
				{4, 4},
			},
			check: true,
		},
		{
			input:    index{2, 1},
			expected: [4]index{},
			check:    false,
		},
		{
			input:    index{1, 2},
			expected: [4]index{},
			check:    false,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			output, check := dummyData.downRight(tt.input)
			assert.Equal(t, tt.expected, output)
			assert.Equal(t, tt.check, check)
		})
	}
}
