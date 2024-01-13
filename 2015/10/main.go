package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	input := "1113122113"
	output := ""
	start := time.Now()

	var iteration time.Time

	for i := 0; i < 50; i++ {

		fmt.Printf("%d/50 - %.2f Minutes (%.2f Minutes) \n", i, time.Since(start).Minutes(), time.Since(iteration).Minutes())
		iteration = time.Now()
		if i != 0 {
			input = strings.Replace(output, "'", "", -1)
			output = ""
		}

		currentRune := ' '
		currentRuneCount := 0
		for i, r := range input {
			if currentRune == ' ' {
				currentRune = r
				currentRuneCount++
			} else if currentRune == r {
				currentRuneCount++
			} else {
				output = fmt.Sprintf("%s%d%q", output, currentRuneCount, currentRune)
				currentRune = r
				currentRuneCount = 1
			}
			if i == len(input)-1 {
				output = fmt.Sprintf("%s%d%q", output, currentRuneCount, r)
			}
		}
	}
	output = strings.Replace(output, "'", "", -1)
	fmt.Println(len(output))
}
