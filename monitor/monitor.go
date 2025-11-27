package monitor

import (
	"log"
	"net/http"
	"time"
)

type Monitoravel interface {
	Verificar() (int, error)
	GetNome() string
}

type Site struct {
	ID     int    `json:"id"`
	Nome   string `json:"nome"`
	URL    string `json:"url"`
	Status int    `json:"status"`
}

var httpClient = &http.Client{
	Timeout: 5 * time.Second,
}

func (s Site) Verificar() (int, error) {
	resp, err := httpClient.Get(s.URL)
	if err != nil {
		s.Status = 0
		return s.Status, err
	}
	defer resp.Body.Close()

	s.Status = resp.StatusCode

	return s.Status, nil
}

func (s Site) GetNome() string {
	return s.Nome
}

func Check(site *Site, store *Store) error {
	status, err := site.Verificar()
	if err != nil {
		log.Printf("Erro ao verificar %s: %v\n", site.URL, err)
	}

	site.Status = status
	err = store.AtualizarStatusSite(site.ID, status)

	return err
}
