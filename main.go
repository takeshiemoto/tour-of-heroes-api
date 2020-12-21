package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
}

func main() {
	mux := httprouter.New()

	const prefix = "/api/v1/"

	mux.GET("/", rootHandler)
	mux.GET(prefix+"heroes", heroesHandler)
	mux.GET(prefix+"heroes/:id", heroHandler)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	server.ListenAndServe()
}

func rootHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var m = struct {
		Message string `json:"message"`
	}{
		Message: "ALIVE",
	}
	bytes, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(bytes)
}

func heroesHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	heroes, err := GetHeroes()
	bytes, err := json.Marshal(heroes)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func heroHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		log.Fatal(err)
	}
	hero, err := GetHeroById(id)

	output, err := json.MarshalIndent(&hero, "", "\t\t")
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
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

func GetHeroById(id int) (hero Hero, err error) {
	hero = Hero{}
	err = Db.QueryRow("SELECT * FROM heroes WHERE id = $1", id).Scan(&hero.ID, &hero.Name, &hero.CreatedAt, &hero.UpdateAt)
	return
}
