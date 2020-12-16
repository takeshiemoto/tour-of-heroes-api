package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=toh dbname=toh password=toh sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/heroes", heroesHandler)

	server.ListenAndServe()
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Tour Of Heroes API")
}

func heroesHandler(w http.ResponseWriter, r *http.Request) {
	heroes, err := GetHeroes()
	bytes, err := json.Marshal(heroes)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func GetHeroes() (heroes Heroes, err error) {
	rows, err := Db.Query("select * from heroes")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		hero := Hero{}
		err = rows.Scan(&hero.ID, &hero.Name, &hero.CreatedAt, &hero.UpdateAt)
		if err != nil {
			return nil, err
		}
		heroes = append(heroes, hero)
	}
	rows.Close()
	return
}
