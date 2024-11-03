package main

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	OPENNESS = iota + 1
	CONSCIENTIOUSNESS
	EXTRAVERSION
	AGREEABLENESS
	NEUROTICISM
)

type Gene struct {
	Character string  // Represents a genetic character
	Value     float32 // Represents the value or weight of the gene
}

type PhysicalTraits struct {
	Strength float32 `json:"strength"`
}

type ConductTraits struct {
	Aggression      float32 `json:"aggression"`
	Cooperation     float32 `json:"cooperation"`
	Competitiveness float32 `json:"competitiveness"`
}

type Traits struct {
	Openness          float32 `json:"openness"`
	Conscientiousness float32 `json:"conscientiousness"`
	Extraversion      float32 `json:"extraversion"`
	Agreeableness     float32 `json:"agreeableness"`
	Neuroticism       float32 `json:"neuroticism"`
}

type Individual struct {
	Genes  []Gene `json:"genes"`
	Traits Traits `json:"traits"`
}

func generateRandomGene() Gene {
	characters := []string{"A", "B", "C", "D", "E"}
	randomChar := characters[rand.Intn(len(characters))]
	return Gene{
		Character: randomChar,
		Value:     rand.Float32(),
	}
}

func generateGenesFromHeritage(genes1, genes2 []Gene) []Gene {
	var genes []Gene
	for i := 0; i < GENES_QUANTITY; i++ {
		if rand.Float32() < 0.5 {
			genes = append(genes, genes1[i])
		} else {
			genes = append(genes, genes2[i])
		}
	}
	return genes
}

func generateIndividual(parents []Being) Individual {
	var individual Individual
	var genes []Gene
	if parents == nil {
		genes := make([]Gene, GENES_QUANTITY)
		for i := 0; i < GENES_QUANTITY; i++ {
			genes[i] = generateRandomGene()
		}
	} else {
		genes = generateGenesFromHeritage(parents[0].genes, parents[1].genes)
	}
	individual.Genes = genes
	individual.Traits = calculateTraits(genes)
	return individual
}

// Calculates traits based on the genes with added randomness
func calculateTraits(genes []Gene) Traits {
	var traits Traits
	for _, gene := range genes {
		switch gene.Character {
		case "A":
			traits.Openness += gene.Value
		case "B":
			traits.Conscientiousness += gene.Value
		case "C":
			traits.Extraversion += gene.Value
		case "D":
			traits.Agreeableness += gene.Value
		case "E":
			traits.Neuroticism += gene.Value
		}
	}
	// Introduce randomness to traits for variability
	traits.Openness += rand.Float32() * 0.1
	traits.Conscientiousness += rand.Float32() * 0.1
	traits.Extraversion += rand.Float32() * 0.1
	traits.Agreeableness += rand.Float32() * 0.1
	traits.Neuroticism += rand.Float32() * 0.1

	fmt.Printf("Traits: %v\n", traits)

	traits.NormalizeTraits()

	return traits
}

func CalculateEuclideanDistance(g1, g2 []Gene) float32 {
	var sum float32
	for i := range g1 {
		diff := g1[i].Value - g2[i].Value
		sum += diff * diff
	}
	return float32(math.Sqrt(float64(sum)))
}

func (traits *Traits) NormalizeTraits() {
	if traits.Openness > 1 {
		traits.Openness = 1
	} else if traits.Openness < 0 {
		traits.Openness = 0
	}
	if traits.Conscientiousness > 1 {
		traits.Conscientiousness = 1
	} else if traits.Conscientiousness < 0 {
		traits.Conscientiousness = 0
	}
	if traits.Extraversion > 1 {
		traits.Extraversion = 1
	} else if traits.Extraversion < 0 {
		traits.Extraversion = 0
	}
	if traits.Agreeableness > 1 {
		traits.Agreeableness = 1
	} else if traits.Agreeableness < 0 {
		traits.Agreeableness = 0
	}
	if traits.Neuroticism > 1 {
		traits.Neuroticism = 1
	} else if traits.Neuroticism < 0 {
		traits.Neuroticism = 0
	}
}
