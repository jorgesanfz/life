package main

import (
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"
)

const (
	lifespan   = 200
	numBeings  = 100
	iterations = 500
	turnPause  = 100 * time.Millisecond
)

var (
	beings     []Being      // Global beings slice
	beingsLock sync.RWMutex // Mutex for safe concurrent access
)

func createBeings() []Being {
	beings = make([]Being, numBeings)
	for i := 0; i < numBeings; i++ {
		beings[i] = *NewBeing(nil)
		beings[i].state()
	}
	return beings
}

func updateBeing(being Being, beingsAux []Being, results chan<- Being) {
	// Update the being and get its status and children
	alive, childs := being.update(beingsAux)
	fmt.Println("Alive:", alive)
	fmt.Println("Childs:", childs)

	// Send new child beings through the channel
	if len(childs) > 0 {
		fmt.Printf("New beings were born!\n")
		for _, child := range childs {
			fmt.Printf("Child %s\n", child.id)
			results <- child // Send each new being through the channel
		}
		fmt.Printf("Being %s sent childs\n", being.id) // Log being update
	}

	// Send the current being back through the channel if it's still alive
	if alive {
		results <- being                            // Send the alive being back through the channel
		fmt.Printf("Being %s is alive\n", being.id) // Log being update
		being.state()                               // Call state method on the alive being
	} else {
		fmt.Printf("Being %s died\n", being.id) // Log if the being has died
	}

	fmt.Printf("Being %s updated\n", being.id) // Log being update
}

func updateBeings(results chan<- Being) {
	var wg sync.WaitGroup
	beingsLock.RLock()         // Acquire a read lock
	defer beingsLock.RUnlock() // Ensure the lock is released when the function exits

	beingsAux := make([]Being, len(beings))
	copy(beingsAux, beings)

	for _, being := range beings {
		wg.Add(1) // Add to WaitGroup before launching goroutine
		go func(b Being) {
			defer wg.Done() // Ensure Done is called
			updateBeing(b, beingsAux, results)
		}(being) // Capture the current being in the closure
	}
	fmt.Println("All beings updated")
	fmt.Println(Red+"TOTAL BEINGS: ", len(beings))

	// Wait for all goroutines to finish
	wg.Wait()
	close(results) // Close the results channel after all goroutines finish
}

func RunSimulation() []Gene {
	var aliveBeings []Being

	for i := 0; i < iterations; i++ {
		fmt.Printf("Iteration %d\n", i)

		results := make(chan Being) // Create a channel to receive results
		go func() {
			for result := range results {
				fmt.Printf("Result %s\n", result.id)
				aliveBeings = append(aliveBeings, result) // Collect alive beings
			}
		}()

		// Update beings concurrently
		updateBeings(results)

		fmt.Println("All beings updated")

		// Use a write lock to safely update the beings slice
		//beingsLock.Lock()
		if len(aliveBeings) >= 50 {
			/*sort.Slice(aliveBeings, func(i, j int) bool {
				return aliveBeings[i].status > aliveBeings[j].status
			})*/
			aliveBeings = aliveBeings[:50] // Limit the number of alive beings
		}
		beings = aliveBeings // Update the beings list
		aliveBeings = nil    // Reset the alive beings list
		//beingsLock.Unlock()

		time.Sleep(turnPause)
	}

	var topGenes []Gene
	sort.Slice(beings, func(i, j int) bool {
		return beings[i].status > beings[j].status
	})

	top10PercentCount := len(beings) / 10
	top10Percent := beings[:top10PercentCount]
	bottom10Percent := beings[len(beings)-top10PercentCount:]

	for _, being := range top10Percent {
		if being.status == 100 {
			topGenes = append(topGenes, being.genes...)
		}
		fmt.Printf("Top Being %s | status: %.2f | genes: %v\n", being.id, being.status, being.genes)
	}

	for _, being := range bottom10Percent {
		fmt.Printf("Bottom Being %s | status: %.2f\n", being.id, being.status)
	}

	//analyse(topGenes)
	return topGenes
}

func RunMultipleSimulations() {
	winners := []Gene{}

	for i := 0; i < iterations; i++ {
		topGenes := RunSimulation()
		winners = append(winners, topGenes...)
		fmt.Printf(Purple+"Simulation %d\n", i)
		//time.Sleep(100 * time.Millisecond)
	}

	fmt.Printf(Cyan+"Winners: %v\n", winners)
	//analyse(winners)
}

func main() {
	createBeings()
	go RunSimulation()

	http.Handle("/beings", CorsMiddleware(http.HandlerFunc(beingsHandler)))
	fmt.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
