package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func lookup(ipToMac map[string]string, macAddress string) string {
	for ip, mac := range ipToMac {
		if strings.Compare(mac, macAddress) == 0 {
			return ip
		}
	}

	return ""
}

func GetIpv4(w http.ResponseWriter, r *http.Request) {
	var requests Ipv4Requests

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&requests)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	log.Println("Requests received:", requests)

	var responses []Ipv4Response

	ipToMac := GetLeases(LeasesFilePath)
	for _, request := range requests {
		responses = append(responses,
			Ipv4Response{
				Mac:  request.Mac,
				Ipv4: lookup(ipToMac, request.Mac),
			})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses)
	log.Println("Responses sent:", responses)
}
