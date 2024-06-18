package main

import (
	"fmt"
	"strconv"
)

func stp() func() {

	return func() {
		if debug {
			fmt.Printf("Stopping the program\n")
		}
		pointer = 999999
	}
}

func hlf(r string) func() {
	var register *uint
	switch r {
	case "a":
		register = &registers.a
	case "b":
		register = &registers.b
	}

	return func() {
		if debug {
			fmt.Printf("Halving register %v\n", r)
		}
		*register = *register / 2
		pointer++
	}
}

func tpl(r string) func() {
	var register *uint
	switch r {
	case "a":
		register = &registers.a
	case "b":
		register = &registers.b
	}
	return func() {
		if debug {
			fmt.Printf("Tripling Register %v\n", r)
		}
		*register = *register * 3
		pointer++
	}
}

func inc(r string) func() {
	var register *uint
	switch r {
	case "a":
		register = &registers.a
	case "b":
		register = &registers.b
	}
	return func() {
		if debug {
			fmt.Printf("Incrementing Register %v\n", r)
		}
		*register++
		pointer++
	}
}

func jmp(o string) func() {
	var j int64
	switch dir := o[0]; dir {
	case '+':
		j, _ = strconv.ParseInt(o[1:], 10, 64)
	case '-':
		j, _ = strconv.ParseInt(o[1:], 10, 64)
		j = -j
	}

	return func() {
		if debug {
			fmt.Printf("Jumping %v\n", o)
		}
		pointer = Pointer(int64(pointer) + j)
	}
}

func jie(r string, o string) func() {
	var register *uint
	switch r {
	case "a":
		register = &registers.a
	case "b":
		register = &registers.b
	}

	var j int64
	switch dir := o[0]; dir {
	case '+':
		j, _ = strconv.ParseInt(o[1:], 10, 64)
	case '-':
		j, _ = strconv.ParseInt(o[1:], 10, 64)
		j = -j
	}

	return func() {
		if int(*register)%2.0 == 0 {
			if debug {
				fmt.Printf("Register is even (%v), Jumping %v\n", *register, o)
			}
			pointer = Pointer(int64(pointer) + j)
		} else {
			if debug {
				fmt.Printf("Register is not even (%v)", *register)
			}
			pointer++
		}

	}
}

func jio(r string, o string) func() {
	var register *uint
	switch r {
	case "a":
		register = &registers.a
	case "b":
		register = &registers.b
	}

	var j int64
	switch dir := o[0]; dir {
	case '+':
		j, _ = strconv.ParseInt(o[1:], 10, 64)
	case '-':
		j, _ = strconv.ParseInt(o[1:], 10, 64)
		j = -j
	}

	return func() {
		if *register == 1 {
			if debug {
				fmt.Printf("Register is one (%v), Jumping %v\n", *register, o)
			}
			pointer = Pointer(int64(pointer) + j)
		} else {
			if debug {
				fmt.Printf("Register is not one (%v)\n", *register)
			}
			pointer++
		}

	}
}
