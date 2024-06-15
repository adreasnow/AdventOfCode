package main

import (
	"errors"
)

func (s *State) spellMagicMissile(playerInstantDamage int) (int, error) {
	if s.PlayerMana > 53 {
		s.PlayerMana -= 53
		s.ManaSpent += 53
		return playerInstantDamage + 4, nil
	} else {
		return 0, errors.New("cannot afford to cast Magic Missile")
	}
}

func (s *State) spellDrain(playerDamage int) (int, error) {
	if s.PlayerMana > 73 {
		s.PlayerMana -= 73
		s.ManaSpent += 73
		s.Player.HP += 2
		return playerDamage + 2, nil
	} else {
		return 0, errors.New("cannot afford to cast Magic Missile")
	}
}

func (s *State) spellShield() error {
	if s.PlayerMana > 113 {
		s.PlayerMana -= 113
		s.ManaSpent += 113
	} else {
		return errors.New("cannot afford to cast Shield")
	}

	if s.TimerShield == 0 {
		s.TimerShield = 6
		return nil
	} else {
		return errors.New("shield already in effect")
	}
}

func (s *State) spellPoison() error {
	if s.PlayerMana > 173 {
		s.PlayerMana -= 173
		s.ManaSpent += 173
	} else {
		return errors.New("cannot afford to cast Poison")
	}

	if s.TimerPoison == 0 {
		s.TimerPoison = 6
		return nil
	} else {
		return errors.New("poison already in effect")
	}
}

func (s *State) spellRecharge() error {
	if s.PlayerMana > 229 {
		s.PlayerMana -= 229
		s.ManaSpent += 229
	} else {
		return errors.New("cannot afford to cast Recharge")
	}

	if s.TimerRecharge == 0 {
		s.TimerRecharge = 5
		return nil
	} else {
		return errors.New("recharge already in effect")
	}
}
