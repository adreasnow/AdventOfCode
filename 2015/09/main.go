package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	"github.com/mowshon/iterium"
)

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

type City int

type Path struct {
	from City
	to   City
}

type Cities map[string]City
type Paths map[Path]int

func AddPath(s1 string, s2 string, d int) {
	var ok bool
	_, ok = cities[s1]
	if !ok {
		cities[s1] = City(len(cities) + 1)
	}
	_, ok = cities[s2]
	if !ok {
		cities[s2] = City(len(cities) + 1)
	}

	paths[Path{cities[s1], cities[s2]}] = d
	paths[Path{cities[s2], cities[s1]}] = d
}

func CalculateDistance(p []string) int {
	distance := 0
	for i := range p {
		if i != len(p)-1 {

			distance += paths[Path{cities[p[i]], cities[p[i+1]]}]
		}
	}
	return distance
}

var paths Paths
var cities Cities

func main() {
	cities = Cities{}
	paths = Paths{}

	strings, err := ReadStrings("input.txt")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, s := range strings {
			var s1, s2 string
			var d int

			fmt.Sscanf(s, "%s to %s = %d", &s1, &s2, &d)
			AddPath(s1, s2, d)
		}

		cityNames := make([]string, 0)
		for s := range cities {
			cityNames = append(cityNames, s)
		}
		permutations, _ := iterium.Permutations(cityNames, len(cityNames)).Slice()

		distances := make([]int, 0)
		for _, p := range permutations {
			distances = append(distances, CalculateDistance(p))
		}
		fmt.Printf("Part A\nThe shortest path is %d\n\n", slices.Min(distances))
		fmt.Printf("Part B\nThe longest path is %d\n", slices.Max(distances))
	}

}
