package main

import (
	"bufio"
	"fmt"
	"os"
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

// func CheckVowels(s string) bool {
// 	vowels := map[rune]int{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1}
// 	sum := 0
// 	for _, char := range s {
// 		sum += vowels[char]
// 	}
// 	return sum >= 3

// }

// func CheckDoubles(s string) bool {
// 	prev := '-'
// 	for _, char := range s {
// 		if char == prev {
// 			return true
// 		} else {
// 			prev = char
// 		}
// 	}
// 	return false
// }

// func CheckBads(s string) bool {
// 	if strings.Contains(s, "ab") || strings.Contains(s, "cd") || strings.Contains(s, "pq") || strings.Contains(s, "xy") {
// 		return false
// 	}
// 	return true
// }

//	func CheckString(s string) bool {
//		if CheckVowels(s) && CheckDoubles(s) && CheckBads(s) {
//			return true
//		} else {
//			return false
//		}
//	}

func CheckRepeatWOOverlap(s string) bool {
	pairs := make(map[string]int, 0)
	prev := '-'
	pair := ""
	for count, char := range s {
		pair = fmt.Sprintf("%v%v", prev, char)
		pos, check := pairs[pair]
		if prev != '-' && check {
			if pos != count-1 {
				return true
			}
		} else {
			pairs[pair] = count
			prev = char
		}
	}
	return false
}

func CheckRepeatWInbetween(s string) bool {
	prevprev := '-'
	prev := '-'
	for _, char := range s {
		if char == prevprev {
			return true
		} else {
			prevprev = prev
			prev = char
		}
	}
	return false
}

func CheckString(s string) bool {
	if CheckRepeatWOOverlap(s) && CheckRepeatWInbetween(s) {
		return true
	} else {
		return false
	}
}

func main() {
	strings, err := ReadStrings("input.txt")
	niceCount := 0
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, s := range strings {
			if CheckString(s) {
				niceCount++
			}
		}
	}

	fmt.Printf("There are %d nice strings out of %d total", niceCount, len(strings))
}
