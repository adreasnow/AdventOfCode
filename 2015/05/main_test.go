package main

import (
	"testing"
)

// func checkVowels(s string) bool {
// func checkDoubles(s string) bool {
// func checkBads(s string) bool {
// func checkString(s string) bool {

// func TestCheckVowels(t *testing.T) {
// 	PassTestCases := []string{"aei",
// 		"xazegov",
// 		"aeiouaeiouaeiou"}
// 	for _, s := range PassTestCases {
// 		if !CheckVowels(s) {
// 			t.Errorf("failed %s", s)
// 		}
// 	}
// 	FailTestCases := []string{"aa",
// 		"aeh",
// 		"ou",
// 		"tydw"}
// 	for _, s := range FailTestCases {
// 		if CheckVowels(s) {
// 			t.Errorf("did not fail %s", s)
// 		}
// 	}
// }

// func TestCheckDoubles(t *testing.T) {
// 	PassTestCases := []string{"xx",
// 		"abcdde",
// 		"aabbccdd"}
// 	for _, s := range PassTestCases {
// 		if !CheckDoubles(s) {
// 			t.Errorf("failed %s", s)
// 		}
// 	}
// 	FailTestCases := []string{"aisdgh",
// 		"aftyuk",
// 		"outksf",
// 		"tydwtyk"}
// 	for _, s := range FailTestCases {
// 		if CheckDoubles(s) {
// 			t.Errorf("did not fail %s", s)
// 		}
// 	}
// }

// func TestCheckBads(t *testing.T) {
// 	PassTestCases := []string{"af",
// 		"ayt",
// 		"ou",
// 		"tydw"}
// 	for _, s := range PassTestCases {
// 		if !CheckBads(s) {
// 			t.Errorf("failed %s", s)
// 		}
// 	}
// 	FailTestCases := []string{"abc",
// 		"cdy",
// 		"xyz",
// 		"opq"}

// 	for _, s := range FailTestCases {
// 		if CheckBads(s) {
// 			t.Errorf("did not fail %s", s)
// 		}
// 	}
// }

// func TestCheckStrings(t *testing.T) {
// 	PassTestCases := []string{"ugknbfddgicrmopn"}
// 	for _, s := range PassTestCases {
// 		if !CheckString(s) {
// 			t.Errorf("failed %s", s)
// 		}
// 	}
// 	FailTestCases := []string{"jchzalrnumimnmhp", // double letter
// 		"haegwjzuvuyypxyu", // contains xy
// 		"dvszwmarrgswjxmb"} // only one vowel
// 	for _, s := range FailTestCases {
// 		if CheckString(s) {
// 			t.Errorf("did not fail %s", s)
// 		}
// 	}
// }

func TestCheckRepeatWOOverlap(t *testing.T) {
	PassTestCases := []string{"xyxy",
		"aabcdefgaa"}
	for _, s := range PassTestCases {
		if !checkRepeatWOOverlap(s) {
			t.Errorf("failed %s", s)
		}
	}
	FailTestCases := []string{"aaa"} // only one vowel
	for _, s := range FailTestCases {
		if checkRepeatWOOverlap(s) {
			t.Errorf("did not fail %s", s)
		}
	}
}

func TestCheckRepeatWInbetween(t *testing.T) {
	PassTestCases := []string{"xyx",
		"abcdefeghi",
		"aaa"}
	for _, s := range PassTestCases {
		if !checkRepeatWInbetween(s) {
			t.Errorf("failed %s", s)
		}
	}
	// FailTestCases := []string{"jchzalrnumimnmhp", // double letter
	// 	"haegwjzuvuyypxyu", // contains xy
	// 	"dvszwmarrgswjxmb"} // only one vowel
	// for _, s := range FailTestCases {
	// 	if CheckRepeatWInbetween(s) {
	// 		t.Errorf("did not fail %s", s)
	// 	}
	// }
}

func TestCheckStrings(t *testing.T) {
	PassTestCases := []string{"qjhvhtzxzqqjkmpb",
		"xxyxx"}
	for _, s := range PassTestCases {
		if !checkString(s) {
			t.Errorf("failed %s", s)
		}
	}
	FailTestCases := []string{"uurcxstgmygtbstg", // no repeat with a single letter between them
		"ieodomkazucvgmuy"} // no pair that appears twice
	for _, s := range FailTestCases {
		if checkString(s) {
			t.Errorf("did not fail %s", s)
		}
	}
}
