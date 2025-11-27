package monitor

import (
	"net/http"
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

func (s Site) Verificar() (int, error) {
	resp, err := http.Get(s.URL)
	if err != nil {
		return 0, err
	}

	s.Status = resp.StatusCode

	return s.Status, nil
}

func (s Site) GetNome() string {
	return s.Nome
}

func Check(site *Site, store *Store) error {
	status, err := site.Verificar()
	if err != nil {
		status = 0
	}

	site.Status = status
	err = store.AtualizarStatusSite(site.ID, status)

	return err
}
