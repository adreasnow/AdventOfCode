package main

import (
	"fmt"
)

type Character struct {
	HP     int
	Damage int
	Armour int
}

type State struct {
	Player          Character
	Boss            Character
	PlayerTurn      bool
	PlayerGoldSpent int
}

type Item struct {
	Cost   int
	Damage int
	Armour int
}

var shop = [][]Item{
	// Weapons
	{
		Item{8, 4, 0},
		Item{10, 5, 0},
		Item{25, 6, 0},
		Item{40, 7, 0},
		Item{74, 8, 0},
	},
	// Armour
	{
		Item{0, 0, 0},
		Item{13, 0, 1},
		Item{31, 0, 2},
		Item{53, 0, 3},
		Item{75, 0, 4},
		Item{102, 0, 5},
	},
	// Rings
	{
		Item{0, 0, 0},
		Item{0, 0, 0},
		Item{25, 1, 0},
		Item{50, 2, 0},
		Item{100, 3, 0},
		Item{20, 0, 1},
		Item{40, 0, 2},
		Item{80, 0, 3},
	},
}

func setupState(armour int, damage int, cost int) State {
	var s State
	s.PlayerTurn = true
	s.Boss = Character{HP: 100, Damage: 8, Armour: 2}
	s.Player = Character{HP: 100, Damage: damage, Armour: armour}
	s.PlayerGoldSpent = cost
	return s
}

func shopPermutations() []State {
	states := make([]State, 0)

	weaponCount := len(shop[0])
	armourCount := len(shop[1])
	ringCount := len(shop[2])

	for weapon := range weaponCount {
		for armour := range armourCount {
			for ring1 := range ringCount {
				for ring2 := range ringCount {
					if ring1 != ring2 {
						armour_stat := shop[1][armour].Armour + shop[2][ring1].Armour + shop[2][ring2].Armour
						damage_stat := shop[0][weapon].Damage + shop[2][ring1].Damage + shop[2][ring2].Damage
						cost := shop[0][weapon].Cost + shop[1][armour].Cost + shop[2][ring1].Cost + shop[2][ring2].Cost
						states = append(states, setupState(armour_stat, damage_stat, cost))
					}
				}
			}
		}
	}
	return states
}

func (s *State) tick() int {
	if s.PlayerTurn {
		s.Boss.HP -= max((s.Player.Damage - s.Boss.Armour), 1)
		s.PlayerTurn = false
	} else {
		s.Player.HP -= max((s.Boss.Damage - s.Player.Armour), 1)
		s.PlayerTurn = true
	}
	return min(s.Player.HP, s.Boss.HP)
}

func main() {
	stateList := shopPermutations()
	minimumGold := 999999
	maximumGold := 0

	for _, state := range stateList {
		lowest_damage := 100
		for lowest_damage > 0 {
			lowest_damage = state.tick()
		}
		if state.Player.HP > 0 {
			if state.PlayerGoldSpent < minimumGold {
				minimumGold = state.PlayerGoldSpent
			}
		} else {
			if state.PlayerGoldSpent > maximumGold {
				maximumGold = state.PlayerGoldSpent
			}
		}
	}
	fmt.Printf("\nMinimum gold spend for win: %d", minimumGold)
	fmt.Printf("\nMaximum gold spend for loss: %d", maximumGold)
}
