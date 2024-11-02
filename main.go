package main

import (
	"fmt"
	"time"
)

const lifespan = 10

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

	return aliveBeings // Return the updated slice with only alive beings
}

func main() {
	beings := createBeings(10)
	for i := 0; i < lifespan; i++ {
		beings = updateBeings(beings)
		time.Sleep(1000)
	}
	for _, being := range beings {
		genes := being.getGenes()
		fmt.Printf("Being %s | status: %f | genes: %v\n", being.id, being.status, genes)
	}
}

/*func RunSimulation() {

}*/
