package main

import (
	"fmt"
	"math"
)

func analyse(winners []Genes) {
	var topAggressionValues, topCooperationValues, statusValues []float32

	fmt.Println("Winners:")
	for i, genes := range winners {
		topAggressionValues = append(topAggressionValues, genes.Aggression)
		topCooperationValues = append(topCooperationValues, genes.Cooperation)
		status := (genes.Aggression + genes.Cooperation) / 2 // Example status calculation
		statusValues = append(statusValues, status)
		fmt.Printf("Simulation %d | Genes: %v | Status: %.2f\n", i, genes, status)
	}

	aggressionMean := mean(topAggressionValues)
	cooperationMean := mean(topCooperationValues)
	statusMean := mean(statusValues)

	aggressionStdDev := stdDev(topAggressionValues, aggressionMean)
	cooperationStdDev := stdDev(topCooperationValues, cooperationMean)
	statusStdDev := stdDev(statusValues, statusMean)

	fmt.Println("\nStatistical Analysis of Genes:")
	fmt.Printf("Aggression - Mean: %.2f, StdDev: %.2f\n", aggressionMean, aggressionStdDev)
	fmt.Printf("Cooperation - Mean: %.2f, StdDev: %.2f\n", cooperationMean, cooperationStdDev)
	fmt.Printf("Status - Mean: %.2f, StdDev: %.2f\n", statusMean, statusStdDev)

	// Calculate and print strategy distribution
	fmt.Println("\nStrategies:")
	strategy(winners)

	// Print Status distribution (percentage above/below average status)
	statusDistribution(statusValues, statusMean)
}

func mean(values []float32) float32 {
	var sum float32
	for _, v := range values {
		sum += v
	}
	return sum / float32(len(values))
}

func stdDev(values []float32, mean float32) float32 {
	var sum float32
	for _, v := range values {
		sum += (v - mean) * (v - mean)
	}
	return float32(math.Sqrt(float64(sum / float32(len(values)))))
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

func statusDistribution(statusValues []float32, mean float32) {
	var above, below int
	for _, status := range statusValues {
		if status >= mean {
			above++
		} else {
			below++
		}
	}
	total := float32(len(statusValues))
	fmt.Println("\nStatus Distribution:")
	fmt.Printf("Above Mean: %.2f%%\n", (float32(above)/total)*100)
	fmt.Printf("Below Mean: %.2f%%\n", (float32(below)/total)*100)
}
