package monitor

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	Db *sql.DB
}

func NewStore() *Store {
	// Abre a conexão com o arquivo global
	db, err := sql.Open("sqlite3", "./monitor.db")
	if err != nil {
		log.Fatal(err)
	}

	stmt := `
  CREATE TABLE IF NOT EXISTS sites (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nome TEXT,
    url TEXT,
    status INTEGER
  );
  `

	_, err = db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}

	return &Store{Db: db}
}

func (s *Store) Adicionar(site Site) error {
	stmt := "INSERT INTO sites(nome, url) VALUES(?, ?)"
	_, err := s.Db.Exec(stmt, site.Nome, site.URL)
	return err
}

func (s *Store) Listar() ([]Site, error) {
	stmt := "SELECT id, nome, url, IFNULL(status, 0) FROM sites"
	rows, err := s.Db.Query(stmt)
	if err != nil {
		fmt.Println("Erro no query")
		fmt.Println(err.Error())
		return nil, err
	}

	defer rows.Close()

	var lista []Site

	for rows.Next() {
		var site Site

		err = rows.Scan(&site.ID, &site.Nome, &site.URL, &site.Status)
		if err != nil {
			fmt.Println("Erro no scan")
			fmt.Println(err.Error())
			return nil, err
		}

		lista = append(lista, site)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return lista, nil
}

func (s *Store) Remover(IDSite int) error {
	stmt := "DELETE FROM sites WHERE id = ?"
	_, err := s.Db.Exec(stmt, IDSite)
	return err
}

func (s *Store) Atualizar(IDSite int, site Site) error {
	stmt := "UPDATE sites SET "

	var args []interface{}
	var partesDaQuery []string

	if site.Nome != "" {
		partesDaQuery = append(partesDaQuery, "nome = ?")
		args = append(args, site.Nome)
	}

	if site.URL != "" {
		partesDaQuery = append(partesDaQuery, "url = ?")
		args = append(args, site.URL)
	}

	camposParaAtualizar := strings.Join(partesDaQuery, ",")

	stmt += camposParaAtualizar + " WHERE id = ?"
	args = append(args, IDSite)

	_, err := s.Db.Exec(stmt, args...)

	return err
}

func (s *Store) AtualizarStatusSite(IDSite int, statusCode int) error {
	stmt := "UPDATE sites SET status = ? WHERE id = ?"

	_, err := s.Db.Exec(stmt, statusCode, IDSite)

	return err
}
