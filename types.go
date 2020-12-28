package main

import (
	"time"
)

type Hero struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

func (h *Hero) Retrieve(id int) (err error) {
	err = Db.QueryRow("SELECT * FROM heroes WHERE id = $1", id).Scan(&h.ID, &h.Name, &h.CreatedAt, &h.UpdateAt)
	return
}

func (h *Hero) Create() (err error) {
	statement := "INSERT INTO heroes (name) VALUES ($1) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(h.Name).Scan(&h.ID)
	return
}

type Heroes struct {
	Heroes []Hero `json:"heroes"`
}

func (h *Heroes) Fetch() (err error) {
	rows, err := Db.Query("SELECT * FROM heroes")
	if err != nil {
		return
	}

	for rows.Next() {
		hero := Hero{}
		err = rows.Scan(&hero.ID, &hero.Name, &hero.CreatedAt, &hero.UpdateAt)
		if err != nil {
			return
		}
		h.Heroes = append(h.Heroes, hero)
	}
	rows.Close()
	return
}
