package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func scanForHex(s string) string {
	escapes := strings.Count(s, "\\")
	for i := 0; i <= escapes; i++ {
		for count, char := range s {
			if count+4 <= len(s) {
				if char == '\\' && s[count+1] == 'x' {
					_, err := strconv.ParseUint(s[count+2:count+4], 16, 16)
					if err == nil {
						replace := string(s[count : count+4])
						s = strings.Replace(s, replace, "U", 1)
					}
				}
			}
		}
	}
	return s
}

func decode(s string) string {
	s = s[1 : len(s)-1]
	s = scanForHex(s)

	s = strings.Replace(s, "\\\\", "|", -1)
	s = strings.Replace(s, "\\\"", "\"", -1)

	return s
}

func encode(s string) string {
	s = strings.Replace(s, "\\", "\\\\", -1)
	s = strings.Replace(s, "\"", "\\\"", -1)

	s = fmt.Sprintf("\"%s\"", s)
	return s
}

func main() {
	strings, err := readStrings("input.txt")
	codeChars := 0
	memChars := 0
	encodedChars := 0

	var decoded string
	var encoded string
	var s string

	if err != nil {
		fmt.Println(err)
	} else {
		for _, s = range strings {
			codeChars += len(s)
			decoded = decode(s)
			memChars += len(decoded)
		}

		for _, s = range strings {
			encoded = encode(s)
			encodedChars += len(encoded)
		}

		for _, s := range []string{"\"\"",
			"\"abc\"",
			"\"aaa\\\"aaa\"",
			"\"\\x27\""} {
			fmt.Println(encode(s))
		}

		fmt.Printf("\nPart A\n----------\nCode: %d\nDecoded: %d\nDifference: %d", codeChars, memChars, codeChars-memChars)
		fmt.Printf("\n\nPart B\n----------\nCode: %d\nEncoded: %d\nDifference: %d", codeChars, encodedChars, encodedChars-codeChars)

	}
}
