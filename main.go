package main

import (
	"fmt"
	"log"
	"sync"

	"health-checker/monitor"
)

func main() {
	store := monitor.NewStore()

	sitesMonitorados, err := store.Listar()
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	canal := make(chan string)

	for _, site := range sitesMonitorados {
		wg.Add(1)
		go monitor.Check(site, &wg, canal)
	}

	go func() {
		wg.Wait()
		close(canal)
	}()

	for msg := range canal {
		fmt.Println(msg)
	}
}
