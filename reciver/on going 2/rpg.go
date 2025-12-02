package main

import "fmt"

type Player struct {
	name              string
	health, maxHealth uint
	mana, maxMana     uint
}

func (player *Player) addHealth(amount uint) {
	player.health += amount
	if player.health > player.maxHealth {
		player.health = player.maxHealth
	}
	fmt.Println(player.name, "add", amount, "health ->", player.health)
}
func (player *Player) applyDamage(amount uint) {
	if player.health-amount > player.health {
		player.health = 0
	} else {
		player.health -= amount
	}
	fmt.Println(player.name, "damage", amount, "health ->", player.health)
}
func (player *Player) addMana(amount uint) {
	player.mana += amount
	if player.mana > player.maxMana {
		player.mana = player.maxMana
	}
	fmt.Println(player.name, "add", amount, "mana ->", player.mana)
}
func (player *Player) consumeMana(amount uint) {
	if player.mana-amount > player.mana {
		player.mana = 0
	} else {
		player.mana -= amount
	}
	fmt.Println(player.name, "damage", amount, "mana ->", player.mana)
}

func main() {
	player := Player{
		name:      "knight",
		health:    100,
		maxHealth: 100,
		maxMana:   500,
		mana:      500,
	}
	player.applyDamage(99)
	player.addHealth(10)
	player.consumeMana(130)
	player.addMana(99)

	player.consumeMana(99999)

}
