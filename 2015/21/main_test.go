package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTick(t *testing.T) {
	assert := assert.New(t)
	var s State
	s.PlayerTurn = true
	s.Boss = Character{HP: 12, Damage: 7, Armour: 2}
	s.Player = Character{HP: 8, Damage: 5, Armour: 5}
	s.tick()
	assert.Equal(s.Boss.HP, 9, "Expecation value")
	s.tick()
	assert.Equal(s.Player.HP, 6, "Expecation value")
	s.tick()
	assert.Equal(s.Boss.HP, 6, "Expecation value")
	s.tick()
	assert.Equal(s.Player.HP, 4, "Expecation value")
	s.tick()
	assert.Equal(s.Boss.HP, 3, "Expecation value")
	s.tick()
	assert.Equal(s.Player.HP, 2, "Expecation value")
	s.tick()
	assert.Equal(s.Boss.HP, 0, "Expecation value")
}
