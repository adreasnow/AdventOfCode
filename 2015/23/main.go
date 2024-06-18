package main

import "fmt"

type Registers struct {
	a uint
	b uint
}

const (
	a Register = "a"
	b          = "b"
)

type Pointer uint
type Register string
type Instructions []func()

func (i *Instructions) add(f func()) {
	*i = append(*i, f)
}

func (i Instructions) run() {
	for int(pointer) <= len(i)-1 {
		i[pointer]()
	}
}

var registers Registers
var pointer Pointer
var debug bool

func main() {
	debug = false
	// debug = true

	registers = Registers{1, 0}
	pointer = 0

	i := Instructions{}

	i.parseInput()
	i.run()
	fmt.Printf("Tape has finished. Register b = %v", registers.b)
}
