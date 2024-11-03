package main

import (
	"fmt"
	"math/rand"
)

type Traits struct {
	Physical PhysicalTraits `json:"physical"`
	Behavior BehaviorTraits `json:"behavior"`
}

type BehaviorTraits struct {
	Aggression  float32 `json:"aggression"`
	Cooperation float32 `json:"cooperation"`
}

type PhysicalTraits struct {
	Strength float32 `json:"strength"`
	Speed    float32 `json:"speed"`
	Size     float32 `json:"size"`
}

// Calculates traits based on the genes with added randomness
func calculateTraits(genes []Gene) Traits {
	var traits Traits
	for _, gene := range genes {
		switch gene.Character {
		case "A":
			traits.Behavior.Aggression += gene.Value
		case "B":
			traits.Behavior.Cooperation += gene.Value
		case "C":
			traits.Physical.Strength += gene.Value
		case "D":
			traits.Physical.Speed += gene.Value
		case "E":
			traits.Physical.Size += gene.Value
		}
	}

	// Add some randomness
	traits.Behavior.Aggression += rand.Float32() * 0.25
	traits.Behavior.Cooperation += rand.Float32() * 0.25
	traits.Physical.Strength += rand.Float32() * 0.25
	traits.Physical.Speed += rand.Float32() * 0.25
	traits.Physical.Size += rand.Float32() * 0.25

	fmt.Printf("Traits: %v\n", traits)

	traits.NormalizeTraits()

	return traits
}

func (traits *Traits) NormalizeTraits() {
	if traits.Physical.Strength > 1 {
		traits.Physical.Strength = 1
	} else if traits.Physical.Strength < 0 {
		traits.Physical.Strength = 0
	}
	if traits.Physical.Speed > 1 {
		traits.Physical.Speed = 1
	} else if traits.Physical.Speed < 0 {
		traits.Physical.Speed = 0
	}
	if traits.Physical.Size > 1 {
		traits.Physical.Size = 1
	} else if traits.Physical.Size < 0 {
		traits.Physical.Size = 0
	}
	if traits.Behavior.Aggression > 1 {
		traits.Behavior.Aggression = 1
	} else if traits.Behavior.Aggression < 0 {
		traits.Behavior.Aggression = 0
	}
	if traits.Behavior.Cooperation > 1 {
		traits.Behavior.Cooperation = 1
	} else if traits.Behavior.Cooperation < 0 {
		traits.Behavior.Cooperation = 0
	}
}
