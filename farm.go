package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/michal-samluk/farm/animals"
	"github.com/michal-samluk/farm/app"
	"log"
	"net/http"
)

func main() {
	context := &app.Context{DB: nil}

	connectDB(context)

	r := mux.NewRouter()
	r.Handle("/animals/", app.Handler{context, animals.IndexHandler}).Methods("GET")
	r.Handle("/animals/", app.Handler{context, animals.CreateHandler}).Methods("POST")
	http.Handle("/", r)

	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func connectDB(c *app.Context) {
	var err error
	c.DB, err = sql.Open("postgres", "user=postgres dbname=farm_development sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}
