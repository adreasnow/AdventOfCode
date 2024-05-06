package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
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

func processStrings(s []string) (map[string][]string, string) {
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

func doSubstitution(replacements map[string][]string, sChan chan Substitution, eChan chan Substitution, ctx context.Context) {
	var kernel string
	var j int
	var replacements_list []string
	var l string
	var mol_len int
	var rep_len int

	for {
		select {
		case m := <-sChan:
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
										eChan <- nm
									} else {
										fmt.Println(nm)
										smallest = len(nm.molecule)
										sChan <- nm
									}
								}
							}
						}
					}
				}
			}
		case <-ctx.Done():
			return
		}
	}
}

func eHandler(eChan chan Substitution, ctx context.Context) {
	for {
		select {
		case e := <-eChan:
			eList = append(eList, e)
			unique["e"] = true
			fmt.Println(e)
		case <-ctx.Done():
			return
		}
	}
}

// func calculateSubstitutions(replacements map[string][]string, molecule string) ([]Substitution, map[string]bool, Substitution) {
func calculateSubstitutions(replacements map[string][]string, molecule string) {
	ctx, cancel := context.WithCancel(context.Background())
	sChan := make(chan Substitution, 99999999999)
	e_chan := make(chan Substitution)
	smallest = len(molecule)

	unique := make(map[string]bool, 0)
	unique["e"] = false
	ue := false
	eList = make([]Substitution, 0)

	sChan <- Substitution{molecule, 0, 0}
	go eHandler(e_chan, ctx)
	for range 1000 {
		go doSubstitution(replacements, sChan, e_chan, ctx)
	}
	for !ue {
		time.Sleep(time.Second)
		ue = unique["e"]
	}
	fmt.Println("Killing Goroutines")
	cancel()
}

// var unique map[string]bool
var eList []Substitution
var unique map[string]bool
var smallest int

func main() {
	input, err := readStrings("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	replacements, molecule := processStrings(input)
	calculateSubstitutions(replacements, molecule)
	fmt.Println(eList)

	// fmt.Printf("It took %d steps to get to 'e'\n", e.changes)
}
