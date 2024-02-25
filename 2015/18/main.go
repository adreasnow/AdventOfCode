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

func (g *Grid) initGrid() {
	for i, l := range g {
		for j := range l {
			(*g)[i][j] = false
		}
	}
}

func (g *Grid) inputToGrid(initial []string) {
	for i, l := range initial {
		for j, state := range l {
			if string(state) == "#" {
				(*g)[i+1][j+1] = true
			} else if string(state) == "." {
				(*g)[i+1][j+1] = false
			}
		}
	}
	gridSize := len(*g) - 2
	(*g)[1][gridSize] = true
	(*g)[gridSize][gridSize] = true
	(*g)[gridSize][1] = true
	(*g)[1][1] = true
}

func (g Grid) countNeighbours(i int, j int) int {
	count := 0
	for _, k := range []int{i - 1, i, i + 1} {
		for _, l := range []int{j - 1, j, j + 1} {
			if g[k][l] {
				count++
			}
		}
	}
	if g[i][j] {
		count--
	}
	return count
}

func (g *Grid) gridStep() {
	gridSize := len(*g)
	var state bool
	var count int
	newG := *g
	for i := range gridSize {
		for j := range gridSize {
			if i != 0 && i != gridSize-1 && j != 0 && j != gridSize-1 {
				count = (*g).countNeighbours(i, j)
				state = (*g)[i][j]
				// The rules
				if !(i == 1 && j == 1) &&
					!(i == 1 && j == gridSize-2) &&
					!(i == gridSize-2 && j == gridSize-2) &&
					!(i == gridSize-2 && j == 1) {

					if state && (count == 2 || count == 3) {
						newG[i][j] = true
					} else {
						newG[i][j] = false
					}
					if !state && count == 3 {
						newG[i][j] = true
					}
				}
			}
		}
	}
	*g = newG
}

func (g Grid) printGrid() {
	var printLine string

	for _, i := range g {
		printLine = ""
		for _, j := range i {
			if j {
				printLine += "#"
			} else {
				printLine += "."
			}
		}
		fmt.Println(printLine)
	}
	fmt.Println("")
}

func (g Grid) countLights() int {
	count := 0
	for i, l := range g {
		for j := range l {
			if g[i][j] {
				count++
			}
		}
	}
	return count
}

type Grid [102][102]bool

func main() {
	initial, err := ReadStrings("input.txt")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	g := Grid{}
	g.initGrid()
	g.inputToGrid(initial)
	g.printGrid()
	for range 100 {
		g.gridStep()
		g.printGrid()
	}
	fmt.Printf("There are %d lights on", g.countLights())
}
