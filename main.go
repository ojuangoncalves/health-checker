package main

import (
	"fmt"
	"log"
	"net/http"

	"health-checker/api"
	"health-checker/monitor"
)

func main() {
	store := monitor.NewStore()
	api := api.API{Store: store}

	port := ":8080"

	http.HandleFunc("/", api.HomeHandler)
	http.HandleFunc("/adicionar", api.CreateHandler)
	http.HandleFunc("/atualizar", api.UpdateHandler)
	http.HandleFunc("/remover", api.DeleteHandler)

	fmt.Printf("Server listening at port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
