package main

import (
	"math"
	"math/rand"
)

type Gene struct {
	Character string  // Represents a genetic character
	Value     float32 // Represents the value or weight of the gene
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
		genes = make([]Gene, GENES_QUANTITY)
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

func CalculateEuclideanDistance(g1, g2 []Gene) float32 {
	if len(g1) != len(g2) {
		//panic("slices g1 and g2 must have the same length")
		return 9
	}

	var sum float32
	for i := range g1 {
		diff := g1[i].Value - g2[i].Value
		sum += diff * diff
	}
	return float32(math.Sqrt(float64(sum)))
}
