package main

import (
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"
)

const (
	lifespan   = 100
	numBeings  = 200
	iterations = 100
	turnPause  = 1000 * time.Millisecond
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

func updateBeings(beings *[]Being) []Being {
	var aliveBeings []Being

	for i := 0; i < len(*beings); i++ {
		beingsLock.Lock()
		being := (*beings)[i]
		alive, _ := being.update(*beings)
		/*if len(childs) > 0 {
			fmt.Print(Green + "New beings were born!\n")
			*beings = append(*beings, childs...)
		}*/
		if alive {
			aliveBeings = append(aliveBeings, being)
			being.state()
		} else {
			fmt.Printf(Red+"Being %s died\n", being.id)
		}
		beingsLock.Unlock()
		time.Sleep(10 * time.Millisecond)
	}

	fmt.Printf("Number of beings alive: %d\n", len(aliveBeings))
	return aliveBeings
}

func RunSimulation() []Genes {
	for i := 0; i < iterations; i++ {
		//beingsLock.Lock()
		beings = updateBeings(&beings)
		//beingsLock.Unlock()
		fmt.Printf(Blue+"Iteration %d\n", i)
		time.Sleep(turnPause)
	}

	var topGenes []Genes
	sort.Slice(beings, func(i, j int) bool {
		return beings[i].status > beings[j].status
	})

	top10PercentCount := len(beings) / 10
	top10Percent := beings[:top10PercentCount]
	bottom10Percent := beings[len(beings)-top10PercentCount:]

	for _, being := range top10Percent {
		if being.status == 100 {
			topGenes = append(topGenes, being.genes)
		}
		fmt.Printf("Top Being %s | status: %.2f | genes: %v\n", being.id, being.status, being.genes)
	}

	for _, being := range bottom10Percent {
		fmt.Printf("Bottom Being %s | status: %.2f\n", being.id, being.status)
	}

	analyse(topGenes)
	return topGenes
}

func RunMultipleSimulations() {
	winners := []Genes{}

	for i := 0; i < iterations; i++ {
		topGenes := RunSimulation()
		winners = append(winners, topGenes...)
		fmt.Printf(Purple+"Simulation %d\n", i)
		//time.Sleep(100 * time.Millisecond)
	}

	analyse(winners)
}

func main() {
	beings = createBeings()
	go RunSimulation()

	http.Handle("/beings", CorsMiddleware(http.HandlerFunc(beingsHandler)))
	fmt.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
