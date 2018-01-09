package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func getValueHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	value, err := s.GetValue(vars["key"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "This address doesnt exist")
		return
	}
	if !strings.HasPrefix(value, "http://") && !strings.HasPrefix(value, "https://") {
		value = "https://" + value
	}
	http.Redirect(w, r, value, http.StatusPermanentRedirect)
}
func createHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.Form.Get("url")
	key, err := s.InsertValue(url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, key)

}
