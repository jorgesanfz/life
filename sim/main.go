package main

import (
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"
)

const (
	lifespan       = 100
	numBeings      = 100
	numSimulations = 100
	turnPause      = 1 * time.Second
)

var (
	beings     []Being
	beingsLock sync.Mutex
)

func createBeings() []Being {
	beings := make([]Being, numBeings)
	for i := 0; i < numBeings; i++ {
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
	for i := 0; i < lifespan; i++ {
		beingsLock.Lock()
		beings = updateBeings(beings)
		beingsLock.Unlock()
		time.Sleep(turnPause)
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

func RunMultipleSimulations() {
	winners := []Genes{}

	for i := 0; i < numSimulations; i++ {
		topGenes := RunSimulation()
		winners = append(winners, topGenes...)
		fmt.Printf(Purple+"Simulation %d\n", i)
		//time.Sleep(100 * time.Millisecond)
	}

	analyse(winners)
}

func main() {
	//RunMultipleSimulations(simulations)
	beings = createBeings()
	go RunSimulation()
	http.Handle("/beings", CorsMiddleware(http.HandlerFunc(beingsHandler)))
	http.ListenAndServe(":8080", nil)
}
