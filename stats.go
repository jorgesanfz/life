package main

import (
	"fmt"
)

func analyse(winners []Genes) {

	var topAggressionValues []float32
	var topCooperationValues []float32

	fmt.Println("Winners:")
	for i, genes := range winners {
		topAggressionValues = append(topAggressionValues, genes.Aggression)
		topCooperationValues = append(topCooperationValues, genes.Cooperation)
		fmt.Printf("Simulation %d | Genes: %v\n", i, genes)
	}

	aggressionMean := mean(topAggressionValues)
	cooperationMean := mean(topCooperationValues)

	fmt.Println("Statistical Analysis of Genes:")
	fmt.Printf("Aggression - Mean: %.2f\n", aggressionMean)
	fmt.Printf("Cooperation - Mean: %.2f\n", cooperationMean)
	fmt.Println("Strategies:")
	strategy(winners)
}

func mean(values []float32) float32 {
	var sum float32
	for _, v := range values {
		sum += v
	}
	return sum / float32(len(values))
}

func strategy(winners []Genes) {
	strat := struct {
		aggression  int
		cooperation int
	}{}
	for _, genes := range winners {
		if genes.Aggression > genes.Cooperation {
			strat.aggression++
		} else {
			strat.cooperation++
		}
	}
	fmt.Printf("Aggression: %d\nCooperation: %d\n", strat.aggression, strat.cooperation)
}
