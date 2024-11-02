package main

import (
	"fmt"
	"sort"
)

const (
	lifespan    = 100
	numBeings   = 100
	simulations = 100
)

func createBeings(n int) []Being {
	beings := make([]Being, n)
	for i := 0; i < n; i++ {
		beings[i] = *NewBeing()
	}
	return beings
}

func updateBeings(beings []Being) []Being {
	var aliveBeings []Being

	for _, being := range beings {
		alive := being.update(beings)
		if alive {
			aliveBeings = append(aliveBeings, being)
			being.state()
		} else {
			fmt.Printf(Red+"Being %s died\n", being.id)
		}
	}

	return aliveBeings
}

func RunSimulation() []Genes {
	beings := createBeings(numBeings)
	for i := 0; i < lifespan; i++ {
		beings = updateBeings(beings)
		//time.Sleep(time.Millisecond)
	}

	sort.Slice(beings, func(i, j int) bool {
		return beings[i].status > beings[j].status
	})

	top10Percent := beings[:len(beings)/10]
	var topGenes []Genes
	for _, being := range top10Percent {
		genes := being.getGenes()
		fmt.Printf("Being %s | status: %f | genes: %v\n", being.id, being.status, genes)
		topGenes = append(topGenes, genes)
	}

	return topGenes
}

func RunMultipleSimulations(n int) {
	winners := []Genes{}

	for i := 0; i < n; i++ {
		topGenes := RunSimulation()
		winners = append(winners, topGenes...)
		fmt.Printf(Purple+"Simulation %d\n", i)
		//time.Sleep(100 * time.Millisecond)
	}

	analyse(winners)
}

func main() {
	RunMultipleSimulations(simulations)
}
