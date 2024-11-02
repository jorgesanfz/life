package main

import "fmt"

type ACTION int

const (
	ATTACK ACTION = iota
	FLEE
	COOPERATE
)

func Interact(a, b *Being) {
	if canInteract(a, b) {
		fmt.Printf(Yellow+"Being %s is interacting with being %s\n", a.id, b.id)
		switch decision(a, b) {
		case 0:
			attack(a, b)
			fmt.Printf(Red+"Being %s is attacking being %s\n", a.id, b.id)
		case 1:
			flee(a, b)
			fmt.Printf(Yellow+"Being %s is fleeing from being %s\n", a.id, b.id)
		case 2:
			cooperate(a, b)
			fmt.Printf(Green+"Being %s is cooperating with being %s\n", a.id, b.id)
		}
	}
}

func canInteract(a, b *Being) bool {
	return a.position.distance(b.position) < INTERACTION_RANGE
}

func decision(a *Being, b *Being) ACTION {
	if a.genes.Aggression > 0.5 && a.strength > b.strength {
		return ATTACK
	} else if a.genes.Cooperation > 0.5 {
		return COOPERATE
	} else {
		return FLEE
	}
}

func attack(a, b *Being) {
	a.status = a.status + (0.1 * b.status)
	b.status = b.status - (0.1 * a.status)
}

func flee(a, b *Being) {
	a.velocity.add(a.position.sub(b.position))
}

func cooperate(a, b *Being) {
	a.status = a.status + (0.05 * b.status)
	b.status = b.status + (0.05 * a.status)
}
