package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextTriangleNumberCoordinate(t *testing.T) {
	tests := []struct {
		ri int
		ci int
		ro int
		co int
	}{
		{
			ri: 1,
			ci: 1,
			ro: 2,
			co: 1,
		},
		{
			ri: 2,
			ci: 1,
			ro: 1,
			co: 2,
		},
		{
			ri: 4,
			ci: 2,
			ro: 3,
			co: 3,
		},
		{
			ri: 1,
			ci: 5,
			ro: 6,
			co: 1,
		},
		{
			ri: 3,
			ci: 1,
			ro: 2,
			co: 2,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d, %d", tt.ri, tt.ci), func(t *testing.T) {
			r, c := nextTriangleNumberCoordinate(tt.ri, tt.ci)
			assert.Equal(t, tt.ro, r)
			assert.Equal(t, tt.co, c)
		})
	}
}

func TestGetValueOfIndex(t *testing.T) {
	value := getValueOfIndex(6, 6)

	assert.Equal(t, 27995004, value)
}

func TestGenerateNextCode(t *testing.T) {

	tests := []struct {
		prev int
		new  int
	}{
		{
			prev: 20151125,
			new:  31916031,
		},
		{
			prev: 31916031,
			new:  18749137,
		},
		{
			prev: 18749137,
			new:  16080970,
		},
		{
			prev: 16080970,
			new:  21629792,
		},
		{
			prev: 21629792,
			new:  17289845,
		},
		{
			prev: 17289845,
			new:  24592653,
		},
		{
			prev: 24592653,
			new:  8057251,
		},
		{
			prev: 8057251,
			new:  16929656,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.prev), func(t *testing.T) {
			new := generateNextCode(tt.prev)
			assert.Equal(t, tt.new, new)
		})
	}
}
