package main

import (
	"testing"
)

func TestDecode(t *testing.T) {
	strings := map[string]int{"\"\"": 0,
		"\"abc\"":        3,
		"\"aaa\\\"aaa\"": 7,
		"\"\\x27\"":      1}

	var decodedLen int
	for s, i := range strings {
		decodedLen = len(Decode(s))
		if decodedLen != i {
			t.Errorf("Decoded %s to len %d instead of %d", s, decodedLen, i)
		}
	}
}

func TestEncode(t *testing.T) {
	strings := map[string]int{"\"\"": 6,
		"\"abc\"":        9,
		"\"aaa\\\"aaa\"": 16,
		"\"\\x27\"":      11}

	var encodedLen int
	for s, i := range strings {
		encodedLen = len(Encode(s))
		if encodedLen != i {
			t.Errorf("Decoded %s to len %d instead of %d", s, encodedLen, i)
		}
	}
}
