package main

import "testing"

func TestNextRune(t *testing.T) {
	newRune, wrap := NextRune('g')
	if newRune != 'h' || wrap != false {
		t.Errorf("Rune of 'g' did not increment to 'h'")
	}

	newRune, wrap = NextRune('z')
	if newRune != 'a' || wrap != true {
		t.Errorf("Rune of 'z' did not wrap to 'a'")
	}
}

func TestNextString(t *testing.T) {
	tests := map[string]string{"abcdefgh": "abcdefgi",
		"tibcdffd": "tibcdffe",
		"zzzzzzzz": "aaaaaaaa",
		"axyzfgjz": "axyzfgka",
		"addezzzz": "addfaaaa",
	}
	for s, expected := range tests {
		if NextString(s) != expected {
			t.Errorf("%s did not match expectation of %v", s, expected)
		}
	}
}

func TestContainsThreeLetterRun(t *testing.T) {
	tests := map[string]bool{"abcdefjh": true,
		"tibcdfft": true,
		"safrycde": true,
		"xyzrtfgj": true,
		"addegtjg": false,
		"hijklmmn": true,
		"hepyyzaa": false,
	}
	for s, expected := range tests {
		if ContainsThreeLetterRun(s) != expected {
			t.Errorf("%s did not match expectation of %v", s, expected)
		}
	}
}

func TestDoesNotContain(t *testing.T) {
	tests := map[string]bool{"abc": true,
		"tibcdff":  false,
		"safcde":   true,
		"xyzofgj":  false,
		"adldegjg": false,
		"hepyyzaa": true,
	}
	for s, expected := range tests {
		if DoesNotContain(s) != expected {
			t.Errorf("%s did not match expectation of %v", s, expected)
		}
	}
}

func TestTwoNonOverlappingPairs(t *testing.T) {
	tests := map[string]bool{"aabbccdd": true,
		"dtibcdff": false,
		"saffcdde": true,
		"xyzofgfj": false,
		"ssadlghh": true,
		"abbceffg": true,
		"abbcegjk": false,
		"abcdeggg": false,
		"ghjaaabc": false,
		"hepyyzaa": true,
	}
	for s, expected := range tests {
		if TwoNonOverlappingPairs(s) != expected {
			t.Errorf("%s did not match expectation of %v", s, expected)
		}
	}
}

func TestCheckString(t *testing.T) {
	tests := map[string]bool{"abcdffaa": true,
		"ghjaabcc": true,
		"hepyyzaa": false,
	}
	for s, expected := range tests {
		if CheckString(s) != expected {
			t.Errorf("%s did not match expectation of %v", s, expected)
		}
	}
}

func TestNextPassword(t *testing.T) {
	tests := map[string]string{"abcdefgh": "abcdffaa",
		"ghijklmn": "ghjaabcc",
	}
	for s, expected := range tests {
		recieved := NextPassword(s)
		if recieved != expected {
			t.Errorf("%s did not match expectation of %v, and instead gave %s", s, expected, recieved)
		}
	}
}
