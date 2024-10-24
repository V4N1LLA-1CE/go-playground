package main

import "testing"

func newPlayer() Player {
	return Player{
		health:    100,
		maxHealth: 100,
		energy:    20,
		maxEnergy: 20,
		name:      "Test",
	}
}

func TestChangeHealth(t *testing.T) {
	p := newPlayer()
	p.changeHealth(999)
	if p.health > p.maxHealth {
		t.Fatalf("Health went over the limit: %v, want: %v", p.health, p.maxHealth)
	}
	p.changeHealth(-(p.maxHealth + 1))
	if p.health < 0 {
		t.Fatalf("Health went below limit: %v, want: %v", p.health, 0)
	}
}
