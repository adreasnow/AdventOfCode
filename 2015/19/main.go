package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
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

func ProcessStrings(s []string) (map[string][]string, string) {
	replacements := make(map[string][]string, 0)
	var key string
	var change string
	var molecule string
	for _, i := range s {
		_, err := fmt.Sscanf(i, "%s => %s ", &change, &key)
		if (err != nil) && (i != "") {
			molecule = i
		} else {
			replacements[key] = append(replacements[key], change)
		}
	}
	return replacements, molecule
}

type Substitution struct {
	molecule string
	position int
	changes  int
}

func DoSubstitution(replacements map[string][]string, s_chan chan Substitution, quit chan int, e_chan chan Substitution) {
	var kernel string
	var j int
	var replacements_list []string
	var l string
	var mol_len int
	var rep_len int

	for {
		select {
		case m := <-s_chan:
			mol_len = len(m.molecule)

			for j = range mol_len { // 0 - n
				for replacement_key := range replacements {
					rep_len = len(replacement_key)

					if j+rep_len <= mol_len {
						kernel = m.molecule[j : j+rep_len]
						replacements_list = replacements[replacement_key]
						if kernel == replacement_key {
							for _, l = range replacements_list {
								substituted := m.molecule[0:j] + l + m.molecule[j+rep_len:]
								nm := Substitution{substituted, j + rep_len - 1, m.changes + 1}
								if len(nm.molecule) <= smallest {
									fmt.Println(len(nm.molecule), smallest)
									if substituted == "e" {
										e_chan <- nm
									} else {
										fmt.Println(nm)
										smallest = len(nm.molecule)
										s_chan <- nm
									}
								}
							}
						}
					}
				}
			}
		case <-quit:
			return
		}
	}
}

func EHandler(e_chan chan Substitution, quit chan int) {
	for {
		select {
		case e := <-e_chan:
			e_list = append(e_list, e)
			unique["e"] = true
			fmt.Println(e)
		case <-quit:
			return
		}
	}
}

// func CalculateSubstitutions(replacements map[string][]string, molecule string) ([]Substitution, map[string]bool, Substitution) {
func CalculateSubstitutions(replacements map[string][]string, molecule string) {
	s_chan := make(chan Substitution, 99999999999)
	quit := make(chan int, 100)
	e_chan := make(chan Substitution)
	smallest = len(molecule)

	unique := make(map[string]bool, 0)
	unique["e"] = false
	ue := false
	e_list = make([]Substitution, 0)

	s_chan <- Substitution{molecule, 0, 0}
	go EHandler(e_chan, quit)
	for range 1000 {
		go DoSubstitution(replacements, s_chan, quit, e_chan)
	}
	for !ue {
		time.Sleep(time.Second)
		ue = unique["e"]
	}
	fmt.Println("Killing Goroutines")
	for range 600 {
		quit <- 0
	}
	// time.Sleep(5 * time.Second)
}

// var unique map[string]bool
var e_list []Substitution
var unique map[string]bool
var smallest int

func main() {
	input, err := ReadStrings("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	replacements, molecule := ProcessStrings(input)
	CalculateSubstitutions(replacements, molecule)
	fmt.Println(e_list)

	// fmt.Printf("It took %d steps to get to 'e'\n", e.changes)
}
