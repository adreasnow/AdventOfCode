package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Aunt struct {
	children    int
	cats        int
	samoyeds    int
	pomeranians int
	akitas      int
	vizslas     int
	goldfish    int
	trees       int
	cars        int
	perfumes    int
}

type Aunts map[int]Aunt

func readStrings(fileName string) ([]string, error) {
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

func (a *Aunts) processAunt(s string) {
	aunt := Aunt{
		children:    -1,
		cats:        -1,
		samoyeds:    -1,
		pomeranians: -1,
		akitas:      -1,
		vizslas:     -1,
		goldfish:    -1,
		trees:       -1,
		cars:        -1,
		perfumes:    -1,
	}

	var n int
	chemicals := map[string]int{}
	var c string
	var q int

	split := strings.Split(s, ":")
	fmt.Sscanf(strings.Split(s, ":")[0], "Sue %d", &n)

	join := strings.Join(split[1:], ":")

	split = strings.Split(join, ",")
	for _, chem := range split {
		chem = strings.Replace(chem, ":", "", -1)
		fmt.Sscanf(chem, "%s %d", &c, &q)
		chemicals[c] = q
	}

	for c, q := range chemicals {
		switch c {
		case "children":
			aunt.children = q
		case "cats":
			aunt.cats = q
		case "samoyeds":
			aunt.samoyeds = q
		case "pomeranians":
			aunt.pomeranians = q
		case "akitas":
			aunt.akitas = q
		case "vizslas":
			aunt.vizslas = q
		case "goldfish":
			aunt.goldfish = q
		case "trees":
			aunt.trees = q
		case "cars":
			aunt.cars = q
		case "perfumes":
			aunt.perfumes = q
		}
	}
	(*a)[n] = aunt
}

func (a1 Aunt) compareAunt(a2 Aunt) bool {
	matches := 0
	if a1.children == a2.children {
		matches++
	}
	if a1.cats < a2.cats {
		matches++
	}
	if a1.samoyeds == a2.samoyeds {
		matches++
	}
	if a1.pomeranians > a2.pomeranians && a2.pomeranians != -1 {
		matches++
	}
	if a1.akitas == a2.akitas {
		matches++
	}
	if a1.vizslas == a2.vizslas {
		matches++
	}
	if a1.goldfish > a2.goldfish && a2.goldfish != -1 {
		matches++
	}
	if a1.trees < a2.trees {
		matches++
	}
	if a1.cars == a2.cars {
		matches++
	}
	if a1.perfumes == a2.perfumes {
		matches++
	}
	if matches == 3 {
		return true
	}
	return false
}

func main() {
	aunts := Aunts{}
	myAunt := Aunt{
		children:    3,
		cats:        7,
		samoyeds:    2,
		pomeranians: 3,
		akitas:      0,
		vizslas:     0,
		goldfish:    5,
		trees:       3,
		cars:        2,
		perfumes:    1,
	}
	input, err := readStrings("input.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, s := range input {
		aunts.processAunt(s)
	}
	for i, a := range aunts {
		if myAunt.compareAunt(a) {
			fmt.Printf("My Aunt sue is Sue #%d", i)
		}
	}
}
