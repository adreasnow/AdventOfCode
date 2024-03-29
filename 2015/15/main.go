package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Ingredient struct {
	capacity   int
	durability int
	flavour    int
	texture    int
	calories   int
}

type Ingredients map[string]Ingredient

type Recipe struct {
	ingredients map[Ingredient]int
	score       int
	calories    int
}

func ReadStrings(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	input := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return input, err
}

func (ingredients *Ingredients) ParseIngredient(s string) {
	var name string
	ingredient := Ingredient{}

	s = strings.Join(strings.Split(s, ":"), "")
	s = strings.Join(strings.Split(s, ","), "")

	fmt.Sscanf(s, "%s capacity %d durability %d flavor %d texture %d calories %d",
		&name,
		&ingredient.capacity,
		&ingredient.durability,
		&ingredient.flavour,
		&ingredient.texture,
		&ingredient.calories)

	(*ingredients)[name] = ingredient

}

func (r *Recipe) CalculateScore() {
	totalCapacity, totalDurability, totalFlavour, totalTexture, totalCalories := 0, 0, 0, 0, 0

	for ingredient, quantity := range r.ingredients {
		totalCapacity += ingredient.capacity * quantity
		totalDurability += ingredient.durability * quantity
		totalFlavour += ingredient.flavour * quantity
		totalTexture += ingredient.texture * quantity
		totalCalories += ingredient.calories * quantity
	}
	for _, property := range []*int{&totalCapacity, &totalDurability, &totalFlavour, &totalTexture, &totalCalories} {
		if *property < 0 {
			*property = 0
		}
	}
	(*r).score = totalCapacity * totalDurability * totalFlavour * totalTexture
	(*r).calories = totalCalories
}

func GeneratePermutations() [][4]int {
	var i, j, k, l int
	perms := make([][4]int, 0)
	for i = 0; i < 100; i++ {
		for j = 0; j < 100; j++ {
			for k = 0; k < 100; k++ {
				for l = 0; l < 100; l++ {
					if i+j+k+l == 100 && i*j*k*l != 0 {
						perms = append(perms, [4]int{i, j, k, l})
					}
				}
			}
		}
	}
	return perms
}

func main() {
	ingredients := Ingredients{}

	input, err := ReadStrings("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, s := range input {
		ingredients.ParseIngredient(s)
	}

	perms := GeneratePermutations()
	i := 0
	maxScore := 0
	_500CalList := make([]Recipe, 0)
	for _, perm := range perms {
		r := Recipe{}
		r.ingredients = make(map[Ingredient]int)

		i = 0
		for _, ingredient := range ingredients {
			r.ingredients[ingredient] = perm[i]
			i++
		}

		r.CalculateScore()
		if r.calories == 500 {
			_500CalList = append(_500CalList, r)
		}
		if maxScore < r.score {
			maxScore = r.score
		}

	}
	fmt.Printf("The highest scoring cookie of any calorie limit is %d.\n", maxScore)

	maxScore = 0
	for _, r := range _500CalList {
		if maxScore < r.score {
			maxScore = r.score
		}
	}
	fmt.Printf("The highest scoring cookie with a 500 calorie limit is %d.\n", maxScore)

}
