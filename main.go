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
	router := httprouter.New()

	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Origin", "*")
			header.Set("Access-Control-Allow-Headers", "Content-Type")
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})

	// index
	router.GET("/", index)

	// Heroes
	router.GET(prefix(HeroesRoute), HeroListHandler)
	router.GET(prefix(HeroRoute), HeroRetrieveHandler)

	router.POST(prefix(HeroesRoute), HeroCreateHandler)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	server.ListenAndServe()
}
