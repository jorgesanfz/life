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
	var child *Being
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
			child = reproduce(a, b)
			fmt.Printf(Pink+"Being %s and being %s are reproducing\n", a.id, b.id)
			fmt.Printf(Green+"Child %s has genes %v\n", child.id, child.genes)
		}
	}
	return child
}

func canInteract(a, b *Being) bool {
	return a.position.distance(b.position) < INTERACTION_RANGE
}

func decision(a *Being, b *Being) ACTION {
	// Calculate genetic distance between two beings
	distance := CalculateEuclideanDistance(a.genes, b.genes)

	// 1. High Genetic Similarity = COOPERATE
	if distance <= GENETIC_SIMILARITY_THRESHOLD {
		// 2. High Status and Genetic Closeness = REPRODUCE
		if distance > 0 && math.Abs(float64(a.status)-float64(b.status)) < STATUS_THRESHOLD {
			fmt.Printf(Orange+"Genetic distance between being %s and being %s is %.2f\n", a.id, b.id, distance)
			fmt.Printf(Orange+"Genetics: being %s: %v, being %s: %v\n", a.id, a.genes, b.id, b.genes)
			fmt.Printf(Orange+"Status: being %s: %.2f, being %s: %.2f\n", a.id, a.status, b.id, b.status)
			return REPRODUCE
		}
		return COOPERATE
	}

	// 3. Low Agreeableness and Low Genetic Similarity = ATTACK
	if a.traits.Behavior.Aggression > AGGRESSION_THRESHOLD && distance > GENETIC_SIMILARITY_THRESHOLD {
		return ATTACK
	}

	// 4. High Cooperation but Low Status or High Distance = FLEE
	if a.traits.Behavior.Cooperation > COOPERATION_THRESHOLD {
		return COOPERATE
	} else {
		return FLEE
	}
}

func attack(a, b *Being) {
	var aChange float32
	var bChange float32
	if a.strength > b.strength {
		aChange = 0.25 * b.status
		bChange = -0.5 * a.status
	} else {
		aChange = -0.5 * b.status
		bChange = 0.25 * a.status
	}
	a.traits.Behavior.Cooperation -= 0.05
	a.traits.NormalizeTraits()
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
	bond(a, b)
	a.updateStatus(0.01 * b.status)
	b.updateStatus(0.01 * a.status)

	child := NewBeing([]Being{*a, *b})
	child.strength = (a.strength + b.strength) / 2
	child.status = (a.status + b.status) / 2
	return child
}

func cooperate(a, b *Being) {
	bond(a, b)
	a.updateStatus(0.01 * b.status)
	b.updateStatus(0.01 * a.status)
}

func bond(a, b *Being) {
	a.traits.Behavior.Cooperation += 0.05
	b.traits.Behavior.Cooperation += 0.025
	a.traits.NormalizeTraits()
	b.traits.NormalizeTraits()
}
