package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/ojuangoncalves/health-checker/monitor"
)

type API struct {
	Store *monitor.Store
}

// Busca todos os sites
func (a *API) HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	sitesMonitorados, err := a.Store.Listar()
	if err != nil {
		http.Error(w, "Erro ao buscar os sites", http.StatusInternalServerError)
		return
	}

	var wg sync.WaitGroup

	for i := range sitesMonitorados {
		wg.Add(1)
		go func(s *monitor.Site) {
			defer wg.Done()
			err = monitor.Check(s, a.Store)
			if err != nil {
				fmt.Println("Erro ao executar o check")
				return
			}
		}(&sitesMonitorados[i])
	}

	wg.Wait()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sitesMonitorados)
}

// Adiciona um site no banco de dados
func (a *API) CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var site monitor.Site
	err := json.NewDecoder(r.Body).Decode(&site)
	if err != nil {
		http.Error(w, "Valor informado inválido", http.StatusBadRequest)
		return
	}

	if (site.Nome == "") || (site.URL == "") {
		http.Error(w, "Informe todos os dados necessários: Nome e URL", http.StatusBadRequest)
		return
	}

	err = a.Store.Adicionar(site)
	if err != nil {
		http.Error(w, "Erro ao adicionar site", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Site adicionado com sucesso"))
}

// Atualiza o nome e/ou URL de um site a partir do ID
func (a *API) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID informado não é válido", http.StatusBadRequest)
		return
	}

	var site monitor.Site
	err = json.NewDecoder(r.Body).Decode(&site)
	if err != nil {
		http.Error(w, "Valor informado não é válido", http.StatusBadRequest)
		return
	}

	err = a.Store.Atualizar(id, site)
	if err != nil {
		http.Error(w, "Erro ao atualizar site", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Site atualizado com sucesso"))
}

// Delete um site a partir do ID
func (a *API) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID informado não é válido", http.StatusBadRequest)
		return
	}

	sites, err := a.Store.Listar()
	if err != nil {
		http.Error(w, "Erro ao buscar sites", http.StatusInternalServerError)
		return
	}

	var idValido bool

	for _, site := range sites {
		if site.ID == id {
			idValido = true
			break
		}
	}

	if !idValido {
		http.Error(w, "ID não encontrado", http.StatusBadRequest)
		return
	}

	err = a.Store.Remover(id)
	if err != nil {
		http.Error(w, "Erro ao remover site", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Site removido com sucesso"))
}
