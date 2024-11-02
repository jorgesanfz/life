package main

import (
	"math/rand"
)

type Genes struct {
	Aggression  float32
	Cooperation float32
	//Speed       float32
	//Strength    float32
	//size	 int
}

func generateRandomGenes() Genes {
	//rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	return Genes{
		Aggression:  rand.Float32(),
		Cooperation: rand.Float32(),
		//Speed:       rand.Float32(),
		//Strength:    rand.Float32(),
	}
}
