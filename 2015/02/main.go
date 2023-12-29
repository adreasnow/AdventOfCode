package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

type present struct {
	l int
	w int
	h int
}

func readPresents(filename string) ([]present, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	presents := make([]present, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		p := present{}
		_, err := fmt.Sscanf(scanner.Text(), "%dx%dx%d", &p.l, &p.w, &p.h)
		if err != nil {
			return nil, err
		}
		presents = append(presents, p)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return presents, err
}

func calcPaper(p *present) int {
	dim := make([]int, 3)
	dim[0] = p.l * p.w
	dim[1] = p.w * p.h
	dim[2] = p.h * p.l
	return (2 * dim[0]) + (2 * dim[1]) + (2 * dim[2]) + slices.Min(dim)
}

func calcRibbon(p *present) int {
	bow := p.l * p.w * p.h

	faces := []int{p.l, p.w, p.h}

	sort.Slice(faces, func(i, j int) bool {
		return faces[i] < faces[j]
	})
	wrap := faces[0] + faces[0] + faces[1] + faces[1]

	return bow + wrap
}

func main() {
	presents, err := readPresents("input.txt")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		paperTotal := 0
		ribbonTotal := 0
		for _, p := range presents {
			paperTotal += calcPaper(&p)
			ribbonTotal += calcRibbon(&p)
		}
		fmt.Printf("%d square feet of paper will be needed\n", paperTotal)
		fmt.Printf("And %d feet of ribbon\n", ribbonTotal)
	}
}
