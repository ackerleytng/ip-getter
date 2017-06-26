package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	flag "github.com/ogier/pflag"
)

var LeasesFilePath string

func main() {
	var listenAddr = flag.String("listen", "0.0.0.0:8000",
		"Address to listen at, such as 0.0.0.0:8000")
	var leasesFilePath = flag.String("path", "./dhcpd.leases.sample",
		"Path to dhcpd.leases")
	flag.Parse()

	LeasesFilePath, _ = filepath.Abs(*leasesFilePath)

	if _, err := os.Stat(LeasesFilePath); err != nil {
		log.Fatal(LeasesFilePath, "does not exist.")
	}

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/ipv4", GetIpv4).Methods("POST")

	log.Println("Server starting...")
	log.Fatal(http.ListenAndServe(*listenAddr, r))
}
