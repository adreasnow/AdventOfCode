package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func SplitListByChar(l []string, delim string) []string {
	split := []string{}
	for _, s := range l {
		out := strings.Split(s, delim)
		split = append(split, out...)
	}
	return split
}

func RemoveRed(s string) string {
	var start int
	out := s
	object := false
	listCloses := 0
	closes := 0
	opens := 0
	replaced := false
	finished := false

	for !finished {
		for i := range out {
			if replaced {
				replaced = false
				break
			}
			if i <= len(out)-3 {
				if replaced {
					break
				}
				if out[i:i+3] == "red" {
					for n := i; n >= 0; n-- { // scan left to find "{" and save the staring index
						if out[n] == '}' {
							closes++
						} else if out[n] == ']' {
							listCloses++
						} else if out[n] == '[' && listCloses > 0 {
							listCloses--
						} else if out[n] == '[' && listCloses == 0 {
							object = false
							closes = 0
							break
						} else if out[n] == '{' && closes > 0 {
							closes--
						} else if out[n] == '{' && closes == 0 {
							start = n
							listCloses = 0
							object = true
							break
						}
					}
					if object {
						for n := i; n <= len(s)-1; n++ { // scan right to find "}" and save the index
							if out[n] == '{' {
								opens++
							} else if out[n] == '}' && opens > 0 {
								opens--
							} else if out[n] == '}' && opens == 0 {
								if out[n+1] == '}' || out[n+1] == ']' { // check if this leaves a trailing comma
									if out[start-1] == ',' {
										start = start - 1
									}
								}
								out = strings.Replace(out, string(out[start:n+1]), "", -1)
								object = false
								replaced = true
								break
							}
						}
					}
				}
			} else {
				finished = true
			}
		}
	}
	return out
}

func main() {
	input, err := ReadStrings("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var out []string
	inputLine := strings.Join(input, "")
	inputLine = RemoveRed(inputLine)
	out = strings.Split(inputLine, ":")
	out = SplitListByChar(out, ",")
	out = SplitListByChar(out, "[")
	out = SplitListByChar(out, "]")
	out = SplitListByChar(out, "{")
	out = SplitListByChar(out, "}")

	sumTotal := 0
	for _, s := range out {
		n, err := strconv.Atoi(s)
		if err == nil {
			sumTotal += n
		}
	}
	fmt.Printf("The sum of all the non-double counted objects is %d", sumTotal)
}
