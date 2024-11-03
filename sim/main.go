package main

import (
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"
)

const (
	lifespan       = 1000
	numBeings      = 100
	numSimulations = 100
	turnPause      = 100 * time.Millisecond
)

var (
	beings     []Being
	beingsLock sync.Mutex
)

func createBeings() []Being {
	beings := make([]Being, numBeings)
	for i := 0; i < numBeings; i++ {
		beings[i] = *NewBeing(Genes{})
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

	fmt.Printf("Number of beings: %d\n", len(aliveBeings))
	return aliveBeings
}

func RunSimulation() []Genes {

	for i := 0; i < lifespan; i++ {
		beingsLock.Lock()
		beings = updateBeings(beings)
		beingsLock.Unlock()
		time.Sleep(turnPause)
	}

	var topGenes []Genes
	top10Percent := beings[:len(beings)/10]
	bottom10Percent := beings[len(beings)-len(beings)/10:]

	sort.Slice(beings, func(i, j int) bool {
		return beings[i].status > beings[j].status
	})

	for _, being := range beings {
		if being.status == 100 {
			topGenes = append(topGenes, being.genes)
		}
	}

	for _, being := range top10Percent {
		genes := being.getGenes()
		fmt.Printf("Being %s | status: %f | genes: %v\n", being.id, being.status, genes)
	}

	sort.Slice(beings, func(i, j int) bool {
		return beings[i].status > beings[j].status
	})

	for _, being := range bottom10Percent {
		fmt.Printf("Being %s | status: %f\n", being.id, being.status)
	}

	analyse(topGenes)

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
	//RunMultipleSimulations()
	beings = createBeings()
	go RunSimulation()
	http.Handle("/beings", CorsMiddleware(http.HandlerFunc(beingsHandler)))
	http.ListenAndServe(":8080", nil)
}
