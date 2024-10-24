package main

import "fmt"

type Player struct {
	health    int
	maxHealth int
	energy    int
	maxEnergy int
	name      string
}

func (p *Player) changeHealth(val int) {

	p.health += val

	if p.health > p.maxHealth {
		p.health = p.maxHealth
	}

	if p.health < 0 {
		p.health = 0
	}

}

func (p *Player) changeEnergy(val int) {

	p.energy += val

	if p.energy > p.maxEnergy {
		p.energy = p.maxEnergy
	}

	if p.energy < 0 {
		p.energy = 0
	}
}

func (p *Player) toString() {
	var status string

	if p.health > 0 {
		status = "(Alive)"
	} else {
		status = "(Dead)"
	}

	fmt.Printf("\nPlayer: %s %s\nHealth: %d\nEnergy: %d\n", p.name, status, p.health, p.energy)
}

func main() {
	p := Player{
		health:    100,
		maxHealth: 100,
		energy:    20,
		maxEnergy: 20,
		name:      "Tarnished",
	}

	p.toString()

	p.changeHealth(20)
	p.changeEnergy(5)
	p.toString()

	p.changeHealth(-24)
	p.changeEnergy(-3)
	p.toString()

	p.changeHealth(-999)
	p.changeEnergy(-90)
	p.toString()

}
