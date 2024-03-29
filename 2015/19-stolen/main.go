package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
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

func ProcessStrings(s []string) ([][]string, string) {
	replacements := make([][]string, 0)
	var key string
	var change string
	var molecule string
	for _, i := range s {
		_, err := fmt.Sscanf(i, "%s => %s ", &key, &change)
		if (err != nil) && (i != "") {
			molecule = i
		} else {
			replacements = append(replacements, []string{key, change})
		}
	}
	return replacements, molecule
}

func GetLongestIndex(l [][]string) ([]string, error) {
	sort.Slice(l, func(i, j int) bool {
		return len(l[i][1]) > len(l[j][1])
	})
	if len(l) == 0 {
		return make([]string, 0), errors.New("l has no length")
	}
	return l[0], nil
}

func RemoveIndex(l [][]string, o []string) [][]string {
	new_l := make([][]string, 0)
	for _, i := range l {
		if i[0] != o[0] || i[1] != o[1] {
			new_l = append(new_l, i)
		}
	}
	return new_l
}

func Calc(molecule string, replacements [][]string) int {
	F := 0
	current := molecule
	for current != "e" {
		f, err := GetLongestIndex(replacements)
		if err != nil {
			replacements = a_replacements
			f, _ = GetLongestIndex(replacements)
		}
		before, after := f[0], f[1]
		new := strings.Replace(current, after, before, 1)
		if current != new {
			F++
		} else {
			replacements = RemoveIndex(replacements, f)
		}
		current = new
	}
	return F
}

var a_replacements [][]string

func main() {
	input, err := ReadStrings("input.txt")
	if err != nil {
		panic(err.Error())
	}
	replacements, molecule := ProcessStrings(input)
	a_replacements = replacements
	count := Calc(molecule, replacements)
	fmt.Printf("It took %d steps to build the molecule.", count)

}