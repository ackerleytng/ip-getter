package main

type Ipv4Request struct {
	Mac string `json:"mac"`
}

type Ipv4Requests []Ipv4Request

type Ipv4Response struct {
	Mac  string `json:"mac"`
	Ipv4 string `json:"ipv4"`
}
