package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func beingsHandler(w http.ResponseWriter, r *http.Request) {
	beingsLock.Lock()
	defer beingsLock.Unlock()

	// Create a copy of the beings slice
	beingsCopy := make([]Being, len(beings))
	copy(beingsCopy, beings)
	fmt.Println(Red+"-----------------Beings copied----------------", len(beingsCopy))

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(beingsCopy); err != nil {
		http.Error(w, "Failed to encode beings", http.StatusInternalServerError)
	}
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
