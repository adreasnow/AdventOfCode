package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckIndexP1(t *testing.T) {
	tests := []struct {
		input    index
		expected int
	}{
		{
			input:    index{1, 1},
			expected: 2,
		},
		{
			input:    index{1, 3},
			expected: 1,
		},
		{
			input:    index{1, 2},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			output := dummyData.checkIndexP1(tt.input)
			assert.Equal(t, tt.expected, output)
		})
	}
}

func TestGenIndexesP1(t *testing.T) {
	tests := []struct {
		input    index
		expected [][4]index
	}{
		{
			input: index{1, 1},
			expected: [][4]index{
				{
					{1, 1},
					{2, 1},
					{3, 1},
					{4, 1},
				},
				{
					{1, 1},
					{1, 2},
					{1, 3},
					{1, 4},
				},
				{
					{1, 1},
					{2, 2},
					{3, 3},
					{4, 4},
				},
			},
		},
		{
			input: index{1, 3},
			expected: [][4]index{
				{
					{1, 3},
					{2, 3},
					{3, 3},
					{4, 3},
				},
				{
					{1, 3},
					{1, 2},
					{1, 1},
					{1, 0},
				},
				{
					{1, 3},
					{2, 2},
					{3, 1},
					{4, 0},
				},
			},
		},
		{
			input: index{3, 1},
			expected: [][4]index{
				{
					{3, 1},
					{2, 1},
					{1, 1},
					{0, 1},
				},
				{
					{3, 1},
					{3, 2},
					{3, 3},
					{3, 4},
				},
				{
					{3, 1},
					{2, 2},
					{1, 3},
					{0, 4},
				},
			},
		},
		{
			input: index{3, 3},
			expected: [][4]index{
				{
					{3, 3},
					{2, 3},
					{1, 3},
					{0, 3},
				},
				{
					{3, 3},
					{3, 2},
					{3, 1},
					{3, 0},
				},
				{
					{3, 3},
					{2, 2},
					{1, 1},
					{0, 0},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			output := dummyData.genIndexesP1(tt.input)
			assert.Len(t, output, len(tt.expected))
			assert.Equal(t, tt.expected, output)
		})
	}
}

func TestGenIndexesP2(t *testing.T) {
	d := data{
		[]rune{'A', 'A', 'A'},
		[]rune{'A', 'X', 'M'},
		[]rune{'A', 'M', 'M'},
	}
	tests := []struct {
		input    index
		expected [4][5]index
	}{
		{
			input: index{1, 1},
			expected: [4][5]index{
				{
					{0, 0},
					{0, 2},
					{1, 1},
					{2, 0},
					{2, 2},
				},
				{
					{2, 0},
					{0, 0},
					{1, 1},
					{2, 2},
					{0, 2},
				},
				{
					{0, 2},
					{2, 2},
					{1, 1},
					{0, 0},
					{2, 0},
				},
				{
					{2, 2},
					{2, 0},
					{1, 1},
					{0, 2},
					{0, 0},
				},
			},
		},
		{
			input:    index{0, 1},
			expected: [4][5]index{},
		},
		{
			input:    index{1, 0},
			expected: [4][5]index{},
		},
		{
			input:    index{2, 0},
			expected: [4][5]index{},
		},
		{
			input:    index{0, 2},
			expected: [4][5]index{},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			output := d.genIndexesP2(tt.input)
			assert.Equal(t, tt.expected, output)
		})
	}
}
