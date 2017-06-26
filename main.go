package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/ipv4", GetIpv4).Methods("POST")

	log.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
