package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	testInput := "ABA\nBAD\nCDE"

	d := parseInput(testInput)

	assert.Equal(t, 'A', (*d)[0][0])
	assert.Equal(t, 'B', (*d)[1][0])
	assert.Equal(t, 'C', (*d)[2][0])

	assert.Equal(t, 'B', (*d)[0][1])
	assert.Equal(t, 'A', (*d)[1][1])
	assert.Equal(t, 'D', (*d)[2][1])

	assert.Equal(t, 'A', (*d)[0][2])
	assert.Equal(t, 'D', (*d)[1][2])
	assert.Equal(t, 'E', (*d)[2][2])

	assert.Len(t, *d, 3)
	assert.Len(t, (*d)[0], 3)
	assert.Len(t, (*d)[1], 3)
	assert.Len(t, (*d)[2], 3)
}

func TestPart1(t *testing.T) {
	input := readInput("input_test_p1.txt")
	d := parseInput(input)
	count := part1(d)
	assert.Equal(t, 18, count)
}

func TestPart2(t *testing.T) {
	input := readInput("input_test_p2.txt")
	d := parseInput(input)
	count := part2(d)
	assert.Equal(t, 9, count)
}
