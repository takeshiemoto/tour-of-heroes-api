package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Ping struct {
	Status int    `json:"status"`
	Result string `json:"result"`
}

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ping := Ping{http.StatusOK, "ok"}
	res, _ := json.Marshal(ping)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Write(res)
}
