package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHlf(t *testing.T) {
	debug = false
	registers = Registers{10, 5}
	pointer = 0

	i := Instructions{}
	i.add(hlf("a"))
	i.add(hlf("b"))

	i.run()

	assert.Equal(t, uint(5), registers.a)
	assert.Equal(t, uint(2), registers.b)
	assert.Equal(t, Pointer(2), pointer)
}

func TestTpl(t *testing.T) {
	debug = false
	registers = Registers{10, 5}
	pointer = 0

	i := Instructions{}
	i.add(tpl("a"))
	i.add(tpl("b"))

	i.run()

	assert.Equal(t, uint(30), registers.a)
	assert.Equal(t, uint(15), registers.b)
	assert.Equal(t, Pointer(2), pointer)
}

func TestInc(t *testing.T) {
	debug = false
	registers = Registers{10, 5}
	pointer = 0

	i := Instructions{}
	i.add(inc("a"))
	i.add(inc("b"))

	i.run()

	assert.Equal(t, uint(11), registers.a)
	assert.Equal(t, uint(6), registers.b)
	assert.Equal(t, Pointer(2), pointer)
}

func TestJmpPlus(t *testing.T) {
	debug = false
	registers = Registers{}
	pointer = 0
	i := Instructions{}

	i.add(inc("a"))
	i.add(jmp("+11"))

	i.run()

	assert.Equal(t, uint(1), registers.a)
	assert.Equal(t, uint(0), registers.b)
	assert.Equal(t, Pointer(12), pointer)
}

func TestJmpMinus(t *testing.T) {
	debug = false
	registers = Registers{}
	pointer = 0
	i := Instructions{}

	i.add(inc("a"))
	i.add(inc("a"))
	i.add(jmp("+2"))
	i.add(stp())
	i.add(inc("a"))
	i.add(jmp("-2"))

	i.run()

	assert.Equal(t, uint(3), registers.a)
	assert.Equal(t, uint(0), registers.b)
	assert.Equal(t, Pointer(999999), pointer)
}

func TestJie(t *testing.T) {
	debug = false
	registers = Registers{}
	pointer = 0
	i := Instructions{}

	i.add(inc("a"))
	i.add(jie("a", "+4"))
	i.add(inc("a"))
	i.add(jie("a", "+2"))

	i.run()

	assert.Equal(t, uint(2), registers.a)
	assert.Equal(t, Pointer(5), pointer)
}

func TestJio(t *testing.T) {
	debug = false
	registers = Registers{}
	pointer = 0
	i := Instructions{}

	i.add(jio("a", "+5"))
	i.add(inc("a"))
	i.add(jio("a", "+3"))
	i.add(inc("a"))
	i.add(inc("a"))

	i.run()

	assert.Equal(t, uint(1), registers.a)
	assert.Equal(t, Pointer(5), pointer)
}
