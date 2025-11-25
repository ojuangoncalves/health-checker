package monitor

import (
	"fmt"
	"net/http"
	"sync"
)

type Monitoravel interface {
	Verificar() (int, error)
	GetNome() string
}

type Site struct {
	ID   int
	Nome string
	URL  string
}

func (s Site) Verificar() (int, error) {
	resp, err := http.Get(s.URL)
	if err != nil {
		return 0, err
	}
	return resp.StatusCode, nil
}

func (s Site) GetNome() string {
	return s.Nome
}

func Check(site Monitoravel, wg *sync.WaitGroup, canal chan string) {
	defer wg.Done()

	status, err := site.Verificar()
	if err != nil {
		canal <- fmt.Sprintf("Site: %s | Erro: %s", site.GetNome(), err.Error())
		return
	}

	canal <- fmt.Sprintf("Site: %s | Status: %d", site.GetNome(), status)
}
