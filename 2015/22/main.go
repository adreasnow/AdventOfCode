package main

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"time"
)

type Character struct {
	HP     int
	Damage int
}

type State struct {
	Player        Character
	Boss          Character
	PlayerTurn    bool
	PlayerMana    int
	ManaSpent     int
	TimerShield   int
	TimerPoison   int
	TimerRecharge int
	Log           []string
}

var spellBook = []string{"MagicMissile", "Drain", "Shield", "Poison", "Recharge"}

func (s *State) LogState() {
	stateString := "State:\n"
	stateString += fmt.Sprintf("\tPlayer HP: %d\n", s.Player.HP)
	stateString += fmt.Sprintf("\tBoss HP: %d\n", s.Boss.HP)
	stateString += fmt.Sprintf("\tPlayer Mana: %d\n", s.PlayerMana)
	stateString += fmt.Sprintf("\tPoison: %d\n", s.TimerPoison)
	stateString += fmt.Sprintf("\tRecharge: %d\n", s.TimerRecharge)
	stateString += fmt.Sprintf("\tShield: %d\n", s.TimerShield)
	s.Log = append(s.Log, stateString)
}

func (s *State) PrintLog() {
	for _, line := range s.Log {
		fmt.Println(line)
	}
}

func setupState() State {
	var s State
	s.PlayerTurn = true
	s.Boss = Character{HP: 71, Damage: 10}
	s.Player = Character{HP: 50}
	s.PlayerMana = 500
	return s
}

type perm []int

func permutations(permLen int, gameChan chan []string) {
	perm := make([]int, permLen+1)
	done := false
	for !done {
		place := 0
		added := false
		for !added {
			if perm[place] < 4 {
				perm[place]++
				added = true
			} else {
				perm[place] = 0
				place++
			}
		}
		if perm[permLen] != 0 {
			done = true
		} else {

			spellList := make([]string, permLen)
			for i, spell := range perm[0:permLen] {
				spellList[i] = spellBook[spell]
			}
			gameChan <- spellList
		}
	}
	return
}

func (s *State) applyConditions() (int, int) {
	playerDamage := 0
	playerArmour := 0

	if s.TimerShield > 0 {
		playerArmour += 7
		s.Log = append(s.Log, "Shield gave 7 armour")
		s.TimerShield--
	}
	if s.TimerPoison > 0 {
		playerDamage += 3
		s.Log = append(s.Log, "Poison dealt 7 damage")
		s.TimerPoison--
	}
	if s.TimerRecharge > 0 {
		s.PlayerMana += 101
		s.Log = append(s.Log, "Recharge returned 101 mana")
		s.TimerRecharge--
	}

	return playerDamage, playerArmour
}

func (s *State) tick(spell string) error {

	playerConditionDamage, playerArmour := s.applyConditions()
	playerDamage := 0

	if s.PlayerTurn {
		var err error

		if hardMode {
			s.Log = append(s.Log, "Player loses 1 HP")
			s.Player.HP--
			if s.Player.HP <= 0 {
				return nil
			}
		}

		switch spell {
		case "MagicMissile":
			playerDamage, err = s.spellMagicMissile(playerDamage)
			s.Log = append(s.Log, "Player cast MagicMissile for 53 mana")
		case "Drain":
			playerDamage, err = s.spellDrain(playerDamage)
			s.Log = append(s.Log, "Player cast Drain for 73 mana")
		case "Shield":
			err = s.spellShield()
			s.Log = append(s.Log, "Player cast Shield for 113 mana")
		case "Poison":
			err = s.spellPoison()
			s.Log = append(s.Log, "Player cast Poison for 173 mana")
		case "Recharge":
			err = s.spellRecharge()
			s.Log = append(s.Log, "Player cast Recharge for 229 mana")
		default:
			return errors.New("Spell not recognised")
		}

		if err != nil {
			return err
		}

		s.Boss.HP -= playerDamage + playerConditionDamage
		s.Log = append(s.Log, fmt.Sprintf("The player did %d damage", playerDamage+playerConditionDamage))
		s.LogState()
		s.PlayerTurn = false
	} else {
		s.Boss.HP -= playerConditionDamage
		if s.Boss.HP <= 0 {
			return nil
		}
		s.Player.HP -= max(s.Boss.Damage-playerArmour, 1)
		s.Log = append(s.Log, fmt.Sprintf("The boss did %d damage", max(s.Boss.Damage-playerArmour, 1)))
		s.LogState()
		s.PlayerTurn = true
	}
	return nil
}

func play(spells []string) (State, bool, int, error) {
	state := setupState()
	var err error
	playerWin := false
	turns := 0

	for _, spell := range spells {
		err = state.tick(spell)
		turns++
		if err != nil {
			return state, false, turns, err
		}

		if state.Player.HP <= 0 {
			break
		}
		if state.Boss.HP <= 0 {
			playerWin = true
			break
		}

		err = state.tick("")
		turns++
		if err != nil {
			return state, false, turns, err
		}

		if state.Player.HP <= 0 {
			break
		}
		if state.Boss.HP <= 0 {
			playerWin = true
			break
		}
	}
	return state, playerWin, turns, nil
}

var hardMode bool

func main() {
	maxTurns := 30
	hardMode = true

	manaList := make([]int, 0)
	gameChan := make(chan []string)
	manaChan := make(chan int)
	runChan := make(chan bool)
	ctx, done := context.WithCancel(context.Background())

	for range 300 {
		go func(gameChan chan []string, manaChan chan int, runChan chan bool, ctx context.Context) {
			for {
				select {
				case game := <-gameChan:
					state, win, turns, err := play(game)

					if err == nil {
						if win {
							fmt.Printf("mana: %d, win: %v, turns: %d\n", state.ManaSpent, win, turns)
							// state.PrintLog()
							manaChan <- state.ManaSpent
							runChan <- true
						}
					}

				case <-ctx.Done():
					return
				}
			}
		}(gameChan, manaChan, runChan, ctx)
	}

	go func(manaChan chan int, ctx context.Context) {
		for {
			select {
			case mana := <-manaChan:
				manaList = append(manaList, mana)
			case <-ctx.Done():
				return
			}
		}
	}(manaChan, ctx)

	go func(runChan chan bool, ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return

			default:
				count := 0
				for range time.Tick(time.Second * 1) {
					fmt.Println("test")

					select {
					case <-runChan:
						count++
					}
				}
				fmt.Printf("Running at %d games/s", count)
			}
		}
	}(runChan, ctx)

	if hardMode {
		fmt.Println("Hard mode is enabled")
	}
	fmt.Println("Running (only wins will be printed)...")
	permutations(maxTurns, gameChan)
	fmt.Println("All permutations have been queued, running for another 10 seconds...")
	time.Sleep(10 * time.Second)
	done()
	fmt.Printf("The lowest mana win game in %d spells is: %d", maxTurns, slices.Min(manaList))
}
