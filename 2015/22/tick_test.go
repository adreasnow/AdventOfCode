package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTick(t *testing.T) {
	var s State
	s.PlayerTurn = true
	s.Boss = Character{HP: 13, Damage: 8}
	s.Player = Character{HP: 10}
	s.PlayerMana = 250

	t.Log("-- Player turn --")
	t.Logf("- Player has %d hit points, %d armor, %d mana", s.Player.HP, 0, s.PlayerMana)
	assert.Equal(t, 10, s.Player.HP)
	assert.Equal(t, 250, s.PlayerMana)
	t.Logf("- Boss has %d hit points", s.Boss.HP)
	assert.Equal(t, 13, s.Boss.HP)
	s.tick("Poison")
	t.Log("Player casts Poison.\n")

	t.Log("-- Boss turn --")
	t.Logf("- Player has %d hit points, %d armor, %d mana", s.Player.HP, 0, s.PlayerMana)
	assert.Equal(t, 10, s.Player.HP)
	assert.Equal(t, 77, s.PlayerMana)
	t.Logf("- Boss has %d hit points", s.Boss.HP)
	assert.Equal(t, 13, s.Boss.HP)
	s.tick("")
	t.Logf("Poison deals 3 damage; its timer is now %d", s.TimerPoison)
	assert.Equal(t, 5, s.TimerPoison)
	t.Logf("Boss attacks for %d damage.\n", s.Boss.Damage)
	assert.Equal(t, 8, s.Boss.Damage)

	t.Log("-- Player turn --")
	t.Logf("- Player has %d hit points, %d armor, %d mana", s.Player.HP, 0, s.PlayerMana)
	assert.Equal(t, 2, s.Player.HP)
	assert.Equal(t, 77, s.PlayerMana)
	t.Logf("- Boss has %d hit points", s.Boss.HP)
	assert.Equal(t, 10, s.Boss.HP)
	s.tick("MagicMissile")
	t.Logf("Poison deals 3 damage; its timer is now %d", s.TimerPoison)
	assert.Equal(t, 4, s.TimerPoison)
	t.Log("Player casts MagicMissile.\n")

	t.Log("-- Boss turn --")
	t.Logf("- Player has %d hit points, %d armor, %d mana", s.Player.HP, 0, s.PlayerMana)
	assert.Equal(t, 2, s.Player.HP)
	assert.Equal(t, 24, s.PlayerMana)
	t.Logf("- Boss has %d hit points", s.Boss.HP)
	assert.Equal(t, 3, s.Boss.HP)
	s.tick("")
	t.Log("Poison deals 3 damage")
	t.Log("This kills the boss, and the player wins")
	assert.Equal(t, 0, s.Boss.HP)
}

func TestTick2(t *testing.T) {
	var s State
	s.PlayerTurn = true
	s.Boss = Character{HP: 14, Damage: 8}
	s.Player = Character{HP: 10}
	s.PlayerMana = 250

	t.Log("-- Player turn --")
	t.Logf("- Player has %d hit points, %d armor, %d mana", s.Player.HP, 0, s.PlayerMana)
	assert.Equal(t, 10, s.Player.HP)
	assert.Equal(t, 250, s.PlayerMana)
	t.Logf("- Boss has %d hit points", s.Boss.HP)
	assert.Equal(t, 14, s.Boss.HP)
	s.tick("Recharge")
	t.Log("Player casts Recharge.\n")

	t.Log("-- Boss turn --")
	t.Logf("- Player has %d hit points, %d armor, %d mana", s.Player.HP, 0, s.PlayerMana)
	assert.Equal(t, 10, s.Player.HP)
	assert.Equal(t, 21, s.PlayerMana)
	t.Logf("- Boss has %d hit points", s.Boss.HP)
	assert.Equal(t, 14, s.Boss.HP)
	s.tick("")
	t.Logf("Recharge provides 101 mana; its timer is now %d.", s.TimerRecharge)
	assert.Equal(t, 4, s.TimerRecharge)
	t.Logf("Boss attacks for %d damage.\n", s.Boss.Damage)
	assert.Equal(t, 8, s.Boss.Damage)

	t.Log("-- Player turn --")
	t.Logf("- Player has %d hit points, %d armor, %d mana", s.Player.HP, 0, s.PlayerMana)
	assert.Equal(t, 2, s.Player.HP)
	assert.Equal(t, 122, s.PlayerMana)
	t.Logf("- Boss has %d hit points", s.Boss.HP)
	assert.Equal(t, 14, s.Boss.HP)
	s.tick("Shield")
	t.Logf("Recharge provides 101 mana; its timer is now %d.", s.TimerRecharge)
	assert.Equal(t, 3, s.TimerRecharge)
	t.Log("Player casts Shield, increasing armor by 7.\n")

	t.Log("-- Boss turn --")
	t.Logf("- Player has %d hit points, %d armor, %d mana", s.Player.HP, 7, s.PlayerMana)
	assert.Equal(t, 2, s.Player.HP)
	assert.Equal(t, 110, s.PlayerMana)
	t.Logf("- Boss has %d hit points", s.Boss.HP)
	assert.Equal(t, 14, s.Boss.HP)
	s.tick("")
	t.Logf("Shield's timer is now %d.", s.TimerShield)
	assert.Equal(t, 5, s.TimerShield)
	t.Logf("Recharge provides 101 mana; its timer is now %d.", s.TimerRecharge)
	assert.Equal(t, 2, s.TimerRecharge)
	t.Log("Boss attacks for 8 - 7 = 1 damage!\n")

	t.Log("-- Player turn --")
	t.Logf("- Player has %d hit point, %d armor, %d mana", s.Player.HP, 7, s.PlayerMana)
	assert.Equal(t, 1, s.Player.HP)
	assert.Equal(t, 211, s.PlayerMana)
	t.Logf("- Boss has %d hit points", s.Boss.HP)
	assert.Equal(t, 14, s.Boss.HP)
	s.tick("Drain")
	t.Logf("Shield's timer is now %d.", s.TimerShield)
	assert.Equal(t, 4, s.TimerShield)
	t.Logf("Recharge provides 101 mana; its timer is now %d.", s.TimerRecharge)
	assert.Equal(t, 1, s.TimerRecharge)
	t.Log("Player casts Drain, dealing 2 damage, and healing 2 hit points..\n")

	t.Log("-- Boss turn --")
	t.Logf("- Player has %d hit points, %d armor, %d mana", s.Player.HP, 7, s.PlayerMana)
	assert.Equal(t, 3, s.Player.HP)
	assert.Equal(t, 239, s.PlayerMana)
	t.Logf("- Boss has %d hit points", s.Boss.HP)
	assert.Equal(t, 12, s.Boss.HP)
	s.tick("")
	t.Logf("Shield's timer is now %d.", s.TimerShield)
	assert.Equal(t, 3, s.TimerShield)
	t.Logf("Recharge provides 101 mana; its timer is now %d.", s.TimerRecharge)
	assert.Equal(t, 0, s.TimerRecharge)
	t.Log("Boss attacks for 8 - 7 = 1 damage!\n")

	t.Log("-- Player turn --")
	t.Logf("- Player has %d hit point, %d armor, %d mana", s.Player.HP, 7, s.PlayerMana)
	assert.Equal(t, 2, s.Player.HP)
	assert.Equal(t, 340, s.PlayerMana)
	t.Logf("- Boss has %d hit points", s.Boss.HP)
	assert.Equal(t, 12, s.Boss.HP)
	s.tick("Poison")
	t.Logf("Shield's timer is now %d.\n", s.TimerShield)
	assert.Equal(t, 2, s.TimerShield)

	t.Log("-- Boss turn --")
	t.Logf("- Player has %d hit points, %d armor, %d mana", s.Player.HP, 7, s.PlayerMana)
	assert.Equal(t, 2, s.Player.HP)
	assert.Equal(t, 167, s.PlayerMana)
	t.Logf("- Boss has %d hit points", s.Boss.HP)
	assert.Equal(t, 12, s.Boss.HP)
	s.tick("")
	t.Logf("Shield's timer is now %d.", s.TimerShield)
	assert.Equal(t, 1, s.TimerShield)
	t.Logf("Poison deals 3 damage; its timer is now %d.", s.TimerPoison)
	assert.Equal(t, 5, s.TimerPoison)
	t.Log("Boss attacks for 8 - 7 = 1 damage!\n")

	t.Log("-- Player turn --")
	t.Logf("- Player has %d hit point, %d armor, %d mana", s.Player.HP, 7, s.PlayerMana)
	assert.Equal(t, 1, s.Player.HP)
	assert.Equal(t, 167, s.PlayerMana)
	t.Logf("- Boss has %d hit points", s.Boss.HP)
	assert.Equal(t, 9, s.Boss.HP)
	s.tick("MagicMissile")
	t.Logf("Shield's timer is now %d.", s.TimerShield)
	assert.Equal(t, 0, s.TimerShield)
	t.Logf("Poison deals 3 damage; its timer is now %d.", s.TimerPoison)
	assert.Equal(t, 4, s.TimerPoison)
	t.Log("Player casts Magic Missile, dealing 4 damage.\n")

	t.Log("-- Boss turn --")
	t.Logf("- Player has %d hit points, %d armor, %d mana", s.Player.HP, 0, s.PlayerMana)
	assert.Equal(t, 1, s.Player.HP)
	assert.Equal(t, 114, s.PlayerMana)
	t.Logf("- Boss has %d hit points", s.Boss.HP)
	assert.Equal(t, 2, s.Boss.HP)
	s.tick("")
	t.Log("Poison deals 3 damage. This kills the boss, and the player wins.\n")
	assert.Equal(t, -1, s.Boss.HP)
}

func playWithState(spells []string, state State) (int, bool, error) {
	var err error
	playerWin := false

	for _, spell := range spells {
		err = state.tick(spell)
		fmt.Print(spell)
		if err != nil {
			return 0, false, err
		}

		if state.Player.HP <= 0 {
			break
		}
		if state.Boss.HP <= 0 {
			playerWin = true
			break
		}

		err = state.tick("")
		if err != nil {
			return 0, false, err
		}

		if state.Player.HP <= 0 {
			break
		}
		if state.Boss.HP <= 0 {
			playerWin = true
			break
		}
	}
	return state.PlayerMana, playerWin, nil
}

func TestPlay(t *testing.T) {
	state := State{}
	state.PlayerTurn = true
	state.Boss = Character{HP: 14, Damage: 8}
	state.Player = Character{HP: 10}
	state.PlayerMana = 250

	mana, win, _ := playWithState([]string{"Recharge", "Shield", "Drain", "Poison", "MagicMissile"}, state)
	assert.Equal(t, 114, mana)
	assert.Equal(t, true, win)

	state = State{}
	state.PlayerTurn = true
	state.Boss = Character{HP: 13, Damage: 8}
	state.Player = Character{HP: 10}
	state.PlayerMana = 250

	mana, win, _ = playWithState([]string{"Poison", "MagicMissile"}, state)
	assert.Equal(t, 24, mana)
	assert.Equal(t, true, win)
}
