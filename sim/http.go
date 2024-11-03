package main

import (
	"encoding/json"
	"net/http"
)

// HTTP Handlers
func beingsHandler(w http.ResponseWriter, r *http.Request) {
	state.mutex.RLock()
	response := struct {
		Generation int     `json:"generation"`
		Beings     []Being `json:"beings"`
	}{
		Generation: state.Generation,
		Beings:     state.Beings,
	}
	state.mutex.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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
