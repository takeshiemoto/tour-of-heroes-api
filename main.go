package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

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

	// index
	mux.GET("/", index)

	// Heroes
	mux.GET(prefix(HeroesRoute), HeroListHandler)
	mux.GET(prefix(HeroRoute), HeroRetrieveHandler)

	mux.POST(prefix(HeroesRoute), HeroCreateHandler)

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
