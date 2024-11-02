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
	REPRODUCE
)

func Interact(a, b *Being) *Being {
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
		case REPRODUCE:
			child := reproduce(a, b)
			fmt.Printf(Green+"Being %s and being %s are reproducing\n", a.id, b.id)
			fmt.Printf(Green+"Child %s has genes %v\n", child.id, child.genes)
			return child
		}
	}
	return nil
}

func canInteract(a, b *Being) bool {
	return a.position.distance(b.position) < INTERACTION_RANGE
}

func decision(a *Being, b *Being) ACTION {
	// Calculate genetic distance between two beings
	distance := CalculateEuclideanDistance(a.genes, b.genes)

	// 1. High Status and Genetic Closeness = REPRODUCE
	if a.status > STATUS_THRESHOLD && b.status > STATUS_THRESHOLD && distance <= GENETIC_SIMILARITY_THRESHOLD {
		return REPRODUCE
	}

	// 2. High Genetic Similarity = COOPERATE
	if distance <= GENETIC_SIMILARITY_THRESHOLD {
		return COOPERATE
	}

	// 3. High Aggression and Low Genetic Similarity = ATTACK
	if a.genes.Aggression > AGGRESSION_THRESHOLD && distance > GENETIC_SIMILARITY_THRESHOLD {
		return ATTACK
	}

	// 4. High Cooperation but Low Status or High Distance = FLEE
	if a.genes.Cooperation > COOPERATION_THRESHOLD {
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

func reproduce(a, b *Being) *Being {
	genes := a.genes.crossover(b.genes)
	child := NewBeing(genes)
	child.strength = (a.strength + b.strength) / 2
	child.status = (a.status + b.status) / 2
	return child
}

func cooperate(a, b *Being) {
	a.updateStatus(0.05 * b.status)
	b.updateStatus(0.05 * a.status)
}
