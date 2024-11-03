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
	iterations = 100000
	turnPause  = 0 * time.Millisecond
	bufferSize = 1000 // Buffer for channels
	workerPool = 200  // Number of workers for processing beings
)

type SimulationState struct {
	Beings     []Being
	Generation int
	mutex      sync.RWMutex
}

type SimulationResult struct {
	Being  Being
	Alive  bool
	Childs []Being
}

// Global simulation state
var state = &SimulationState{
	Beings: make([]Being, 0, numBeings),
}

// Worker pool for processing beings
func beingWorker(id int, jobs <-chan Being, results chan<- SimulationResult, beingsSnapshot []Being, wg *sync.WaitGroup) {
	defer wg.Done()
	for being := range jobs {
		alive, childs := being.update(beingsSnapshot)
		results <- SimulationResult{
			Being:  being,
			Alive:  alive,
			Childs: childs,
		}
	}
}

func createBeings() {
	state.mutex.Lock()
	defer state.mutex.Unlock()

	state.Beings = make([]Being, numBeings)
	for i := 0; i < numBeings; i++ {
		state.Beings[i] = *NewBeing(nil)
		state.Beings[i].state()
	}
}

func processSimulationResults(results <-chan SimulationResult) []Being {
	aliveBeings := make([]Being, 0, numBeings*2) // Pre-allocate with room for growth

	// Process results as they come in
	for result := range results {
		if result.Alive {
			aliveBeings = append(aliveBeings, result.Being)
		}
		aliveBeings = append(aliveBeings, result.Childs...)
	}

	// Sort and trim population if needed
	if len(aliveBeings) > 200 {
		sort.Slice(aliveBeings, func(i, j int) bool {
			return aliveBeings[i].status > aliveBeings[j].status
		})
		aliveBeings = aliveBeings[:200]
	}

	return aliveBeings
}

func updateBeings() []Being {
	// Create buffered channels for better performance
	jobs := make(chan Being, bufferSize)
	results := make(chan SimulationResult, bufferSize)

	// Create a snapshot of current beings for concurrent access
	state.mutex.RLock()
	beingsSnapshot := make([]Being, len(state.Beings))
	copy(beingsSnapshot, state.Beings)
	state.mutex.RUnlock()

	// Start worker pool
	var wg sync.WaitGroup
	for i := 0; i < workerPool; i++ {
		wg.Add(1)
		go beingWorker(i, jobs, results, beingsSnapshot, &wg)
	}

	// Send jobs to workers
	go func() {
		for _, being := range beingsSnapshot {
			jobs <- being
		}
		close(jobs)
	}()

	// Wait for all workers to finish in a separate goroutine
	go func() {
		wg.Wait()
		close(results)
	}()

	return processSimulationResults(results)
}

func RunSimulation() []Gene {
	state.Generation = 0
	topGenes := make([]Gene, 0)

	for i := 0; i < iterations; i++ {
		state.Generation = i

		// Update beings and get new population
		aliveBeings := updateBeings()

		// Update global state
		state.mutex.Lock()
		state.Beings = aliveBeings
		state.mutex.Unlock()

		time.Sleep(turnPause)
	}

	// Process final results
	state.mutex.RLock()
	defer state.mutex.RUnlock()

	sort.Slice(state.Beings, func(i, j int) bool {
		return state.Beings[i].status > state.Beings[j].status
	})

	// Get top 10%
	top10PercentCount := len(state.Beings) / 10
	top10Percent := state.Beings[:top10PercentCount]

	for _, being := range top10Percent {
		if being.status == 100 {
			topGenes = append(topGenes, being.genes...)
		}
	}

	fmt.Println("Top genes:" + fmt.Sprint(topGenes))

	return topGenes
}

func main() {
	// Create initial population
	createBeings()

	// Set up HTTP server with improved routing
	mux := http.NewServeMux()
	mux.Handle("/beings", CorsMiddleware(http.HandlerFunc(beingsHandler)))

	// Start simulation in background
	go RunSimulation()

	// Configure server with timeouts
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Println("Server starting on :8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
