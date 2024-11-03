package main

import (
	"math"
	"math/rand"
	"time"
)

type Genes struct {
	Aggression  float32 `json:"aggression"`
	Cooperation float32 `json:"cooperation"`
	// Speed       float32
	// Strength    float32
}

func generateRandomGenes() Genes {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	return Genes{
		Aggression:  rand.Float32(),
		Cooperation: rand.Float32(),
		// Speed:       rand.Float32(),
		// Strength:    rand.Float32(),
	}
}

func (g Genes) crossover(other Genes) Genes {
	return Genes{
		Aggression:  (g.Aggression + other.Aggression) / 2,
		Cooperation: (g.Cooperation + other.Cooperation) / 2,
		// Speed:       (g.Speed + other.Speed) / 2,
		// Strength:    (g.Strength + other.Strength) / 2,
	}
}

// CalculateEuclideanDistance calculates the Euclidean distance between two Genes
func CalculateEuclideanDistance(g1, g2 Genes) float32 {
	aggressionDiff := g1.Aggression - g2.Aggression
	cooperationDiff := g1.Cooperation - g2.Cooperation
	// Add more differences if you include additional traits

	return float32(math.Sqrt(float64(aggressionDiff*aggressionDiff + cooperationDiff*cooperationDiff)))
}

// ShouldCollaborate determines if two individuals should collaborate based on genetic closeness
/*func ShouldCollaborate(g1, g2 Genes, threshold float32) bool {
	return CalculateEuclideanDistance(g1, g2) <= COLLABORATION_THRESHOLD
}*/
