package main

import (
	"github.com/gorilla/mux"
	"net/http"

	"log"

)

var s StoragePG

func main() {

	err := s.ConnectTo("127.0.0.1", 5432, "smol", "postgres", "postgres")
	if (err != nil) {
		println(err.Error())
		return
	}
	r := mux.NewRouter()

	r.HandleFunc("/create", createHandler).Methods("POST")
	r.HandleFunc("/link/{key}", getValueHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))

}
