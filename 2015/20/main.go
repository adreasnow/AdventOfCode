package main

import (
	"fmt"
	"sync"
)

type Elves struct {
	elves map[int]int
	mu    sync.Mutex
}

type House struct {
	house    int
	presents int
}

func (e *Elves) Increment(elf int) {
	e.mu.Lock()
	e.elves[elf]++
	e.mu.Unlock()
}

func (e *Elves) Get(elf int) int {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.elves[elf]
}

func calculatePresents(nChan chan int, houseChan chan House, quit chan int, elves *Elves) {
	for {
		select {
		case n := <-nChan:
			presents := 0
			for elf := 1; elf <= n/2; elf++ {
				if n%elf == 0 && elves.Get(elf) < 50 {
					presents += elf * 11
					elves.Increment(elf)
				}
			}
			presents += n * 11
			elves.Increment(n)

			houseChan <- House{n, presents}
		case <-quit:
			return
		}
	}
}

func presentsChecker(houseChan chan House, quit chan int, complete *bool, input int) {
	for {
		house := <-houseChan
		fmt.Println(house.house, ": ", house.presents)
		if house.presents >= input {
			*complete = true
			fmt.Printf("The first house to get %d presents is %d", input, house.house)
			for range 20 {
				quit <- 1
				return
			}
		}
	}
}

func main() {
	input := 29000000
	house := 0

	elves := Elves{}
	elves.elves = make(map[int]int, 0)

	n_chan := make(chan int, 100)
	house_chan := make(chan House, 100)
	quit := make(chan int, 100)

	var complete bool

	go presentsChecker(house_chan, quit, &complete, input)

	for range 20 {
		go calculatePresents(n_chan, house_chan, quit, &elves)
	}

	for !complete {
		house++
		n_chan <- house
	}
}
