package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ojuangoncalves/health-checker/api"
	"github.com/ojuangoncalves/health-checker/monitor"
)

func main() {
	store := monitor.NewStore()
	defer store.Db.Close()

	api := api.API{Store: store}

	mux := http.NewServeMux()
	mux.HandleFunc("/", api.HomeHandler)
	mux.HandleFunc("/adicionar", api.CreateHandler)
	mux.HandleFunc("/atualizar", api.UpdateHandler)
	mux.HandleFunc("/remover", api.DeleteHandler)

	handler := corsMiddleware(mux)

	port := ":8080"
	fmt.Printf("Server listening at port %s\n", port)
	log.Fatal(http.ListenAndServe(port, handler))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
