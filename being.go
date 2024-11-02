package main

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

type Being struct {
	id       uuid.UUID
	genes    Genes
	position Vector
	velocity Vector
	age      int
	status   float32
	strength float32
	//acceleration Acceleration
	//health   int
	// maxSpeed   int
}

func NewBeing() *Being {
	return &Being{
		id:       uuid.New(),
		position: Vector{X: rand.Float64(), Y: rand.Float64()},
		velocity: Vector{X: 0, Y: 0},
		status:   20,
		genes:    generateRandomGenes(),
		strength: rand.Float32(),
		// maxSpeed:   1,
	}
}

func (b *Being) update(beings []Being) bool {
	b.age++
	if b.age > lifespan || b.status < 5 {
		return false
	}
	b.move()
	for _, other := range beings {
		Interact(b, &other)
	}
	return true
}

func (b *Being) state() {
	fmt.Printf(Cyan+"Being %s: position: %v, velocity: %v, status: %v, age: %v\n", b.id, b.position, b.velocity, b.status, b.age)
}

func (b *Being) getGenes() Genes {
	return b.genes
}

func (b *Being) move() {
	// Adjust velocity based on boundary conditions
	if b.position.X <= 0 {
		b.velocity.X *= -0.05 // Reverse and dampen velocity when hitting left boundary
		b.position.X = 0      // Clamp position to the left boundary
	} else if b.position.X >= 1 {
		b.velocity.X *= -0.05 // Reverse and dampen velocity when hitting right boundary
		b.position.X = 1      // Clamp position to the right boundary
	} else {
		b.velocity.X += rand.Float64()*0.1 - 0.05 // Random acceleration
	}

	if b.position.Y <= 0 {
		b.velocity.Y *= -0.05 // Reverse and dampen velocity when hitting bottom boundary
		b.position.Y = 0      // Clamp position to the bottom boundary
	} else if b.position.Y >= 1 {
		b.velocity.Y *= -0.05 // Reverse and dampen velocity when hitting top boundary
		b.position.Y = 1      // Clamp position to the top boundary
	} else {
		b.velocity.Y += rand.Float64()*0.1 - 0.05 // Random acceleration
	}

	// Update position based on velocity
	b.position.add(b.velocity)

	// Optional: Ensure position stays within bounds after moving
	b.position.X = clamp(b.position.X, 0, 1)
	b.position.Y = clamp(b.position.Y, 0, 1)
}

// Helper function to clamp a value between a minimum and maximum
func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
