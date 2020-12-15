package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

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
	heroes := Heroes{
		{ID: 1, Name: "John", CreatedAt: time.Now(), UpdateAt: time.Now()},
		{ID: 2, Name: "Paul", CreatedAt: time.Now(), UpdateAt: time.Now()},
		{ID: 3, Name: "George", CreatedAt: time.Now(), UpdateAt: time.Now()},
		{ID: 4, Name: "Ringo", CreatedAt: time.Now(), UpdateAt: time.Now()},
	}
	bytes, err := json.Marshal(heroes)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
