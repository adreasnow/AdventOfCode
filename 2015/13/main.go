package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

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

func ParseRule(rule string, rules *map[string]int, attendees *[]string) {
	var attendeeName string
	var partnerName string
	var happiness int

	rule = strings.Split(rule, ".")[0]
	if strings.Contains(rule, "gain") {
		fmt.Sscanf(rule, "%s would gain %d happiness units by sitting next to %s.", &attendeeName, &happiness, &partnerName)
	} else if strings.Contains(rule, "lose") {
		fmt.Sscanf(rule, "%s would lose %d happiness units by sitting next to %s.", &attendeeName, &happiness, &partnerName)
		happiness = -happiness
	}

	if !slices.Contains(*attendees, attendeeName) {
		*attendees = append(*attendees, attendeeName)
	}

	key := fmt.Sprintf("%s-%s", attendeeName, partnerName)

	(*rules)[key] = happiness
}

func CalculateHappiness(seating *[]string, rules *map[string]int) int {
	var key1 string
	var key2 string

	totalHappiness := 0

	for i := range *seating {
		if i == len(*seating)-1 {
			key1 = fmt.Sprintf("%s-%s", (*seating)[i], (*seating)[0])
			key2 = fmt.Sprintf("%s-%s", (*seating)[0], (*seating)[i])
		} else {
			key1 = fmt.Sprintf("%s-%s", (*seating)[i], (*seating)[i+1])
			key2 = fmt.Sprintf("%s-%s", (*seating)[i+1], (*seating)[i])
		}
		totalHappiness += (*rules)[key1]
		totalHappiness += (*rules)[key2]
	}
	return totalHappiness
}

func main() {
	input, err := ReadStrings("input-2.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	rules := make(map[string]int)
	attendees := make([]string, 0)

	for _, s := range input {
		ParseRule(s, &rules, &attendees)
	}

	permutations, _ := iterium.Permutations(attendees, len(attendees)).Slice()

	happiness := make([]int, 0)

	for _, seating := range permutations {
		happiness = append(happiness, CalculateHappiness(&seating, &rules))
	}

	fmt.Printf("Max happiness is %d\n", slices.Max(happiness))
}
