package main

import (
	"fmt"
	"strings"
)

func NextRune(i rune) (rune, bool) {
	var out int
	var wrap bool
	if i <= 'y' {
		out = int(i) + 1
		wrap = false
	} else if i == 'z' {
		out = 'a'
		wrap = true
	}
	return rune(out), wrap
}

func NextString(s string) string {
	var newRune rune
	wrap := true
	i := len(s) - 1
	runeList := []rune(s)
	for wrap {
		if i == -1 {
			return "aaaaaaaa"
		}
		newRune, wrap = NextRune(runeList[i])
		runeList[i] = newRune
		i--
	}
	return string(runeList)
}

func ContainsThreeLetterRun(s string) bool {
	var i2 rune
	var i3 rune
	var wrap2 bool
	var wrap3 bool

	for i := 0; i <= len(s)-3; i++ {
		i2, wrap2 = NextRune(rune(s[i]))
		i3, wrap3 = NextRune(rune(s[i+1]))
		if i2 == rune(s[i+1]) && i3 == rune(s[i+2]) && !(wrap2 || wrap3) {
			return true
		}
	}
	return false
}

func DoesNotContain(s string) bool {
	return !(strings.Contains(s, "i") || strings.Contains(s, "l") || strings.Contains(s, "o"))
}

func TwoNonOverlappingPairs(s string) bool {
	overlapCount := 0
	previousRune := rune(10)
	previousIndex := -10

	for i := 0; i <= len(s)-2; i++ {
		if rune(s[i]) == rune(s[i+1]) {
			if previousRune != rune(s[i]) && previousIndex != i-1 {
				overlapCount++
				previousRune = rune(s[i])
				previousIndex = i
			}
		}
	}
	return overlapCount >= 2
}

func CheckString(s string) bool {
	return ContainsThreeLetterRun(s) && DoesNotContain(s) && TwoNonOverlappingPairs(s)
}

func NextPassword(s string) string {
	found := false
	for !found {
		s = NextString(s)
		found = CheckString(s)
	}
	return s
}

func main() {
	input := "hepxcrrq"
	fmt.Printf("Part 1: Santa's next password is \"%s\"\n", NextPassword(input))
	fmt.Printf("Part 2: Santa's next password is \"%s\"", NextPassword(NextPassword(input)))
}
