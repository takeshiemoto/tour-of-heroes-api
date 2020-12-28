package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

const (
	HeroesRoute = "/heroes"
	HeroRoute   = "/heroes/:id"
)

func HeroListHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var heroes Heroes
	err := heroes.Fetch()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	bytes, err := json.Marshal(heroes.Heroes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func HeroRetrieveHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var h Hero
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.Retrieve(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(&h, "", "\t\t")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func HeroCreateHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var hero Hero
	err := json.Unmarshal(body, &hero)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = hero.Create()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)
	return
}
