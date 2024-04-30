package main

import "testing"

func TestCalculateScore(t *testing.T) {
	butterscotch := Ingredient{capacity: -1, durability: -2, flavour: 6, texture: 3, calories: 8}
	cinnamon := Ingredient{capacity: 2, durability: 3, flavour: -2, texture: -1, calories: 3}

	r := Recipe{ingredients: map[Ingredient]int{butterscotch: 44, cinnamon: 56}}
	r.calculateScore()
	if r.score != 62842880 {
		t.Errorf("recipe did not match expectation of %d, and instead gave %d", 62842880, r.score)
	}
}
