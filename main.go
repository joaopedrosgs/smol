package smol

import (
	"github.com/gorilla/mux"
	"net/http"

	"log"
)

var s storagePG

func RunSmol(address string, port int, dbname string, username string, password string, ssl string, host_port string) {

	err := s.ConnectTo(address, port, dbname, username, password, ssl)
	if (err != nil) {
		println(err.Error())
		return
	}
	r := mux.NewRouter()

	r.HandleFunc("/create", createHandler).Methods("POST")
	r.HandleFunc("/link/{key}", getValueHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.Handle("/", r)
	err = http.ListenAndServe(":"+host_port, r)
	if (err != nil) {
		log.Println("Smol failed to listen to: " + host_port)
		log.Fatal(err.Error())
	} else {
		log.Println("Smol listening to port: " + host_port)
	}
}
