package main

import (
	"context"
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

func calculatePresents(nChan chan int, houseChan chan House, elves *Elves, ctx context.Context) {
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
		case <-ctx.Done():
			return
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

	var complete bool

	ctx, cancel := context.WithCancel(context.Background())

	go func(houseChan chan House, complete *bool, input int, cancel context.CancelFunc) {
		for {
			house := <-houseChan
			fmt.Println(house.house, ": ", house.presents)
			if house.presents >= input {
				cancel()
				*complete = true
				fmt.Printf("The first house to get %d presents is %d\n", input, house.house)

			}
		}
	}(house_chan, &complete, input, cancel)

	for range 20 {
		go calculatePresents(n_chan, house_chan, &elves, ctx)
	}

	for !complete {
		house++
		n_chan <- house
	}
}
