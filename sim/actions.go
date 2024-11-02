package main

import (
	"fmt"
	"math"
)

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
		case ATTACK:
			attack(a, b)
			fmt.Printf(Red+"Being %s is attacking being %s\n", a.id, b.id)
		case FLEE:
			flee(a, b)
			fmt.Printf(Yellow+"Being %s is fleeing from being %s\n", a.id, b.id)
		case COOPERATE:
			cooperate(a, b)
			fmt.Printf(Green+"Being %s is cooperating with being %s\n", a.id, b.id)
		}
	}
}

func canInteract(a, b *Being) bool {
	return a.position.distance(b.position) < INTERACTION_RANGE
}

func decision(a *Being, b *Being) ACTION {
	if a.genes.Aggression > 0.5 {
		return ATTACK
	} else if a.genes.Cooperation > 0.5 {
		return COOPERATE
	} else {
		return FLEE
	}
}

func attack(a, b *Being) {
	var aChange float32
	var bChange float32
	if a.strength > b.strength {
		aChange = 0.1 * b.status
		bChange = -0.1 * a.status
	} else {
		aChange = -0.1 * b.status
		bChange = 0.1 * a.status
	}
	a.updateStatus(aChange)
	b.updateStatus(bChange)
}

func flee(a, b *Being) {
	direction := a.position.sub(b.position)
	distance := math.Sqrt(direction.X*direction.X + direction.Y*direction.Y)
	if distance > 0 {
		direction.X /= distance
		direction.Y /= distance
	}
	fleeSpeed := 0.1
	a.velocity.X += direction.X * fleeSpeed
	a.velocity.Y += direction.Y * fleeSpeed
}

func cooperate(a, b *Being) {
	a.updateStatus(0.05 * b.status)
	b.updateStatus(0.05 * a.status)
}
