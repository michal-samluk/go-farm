package main

import (
	"github.com/gorilla/mux"
	"github.com/michal-samluk/farm/animals"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/animals/", animals.IndexHandler).Methods("GET")
	r.HandleFunc("/animals/", animals.CreateHandler).Methods("POST")
	http.Handle("/", r)

	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
