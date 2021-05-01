package main

import (
	"khoros-idea-connector/plumbing"
	"log"
	"net/http"
)

func main() {
	for _, r := range plumbing.Routes() {
		http.HandleFunc(r.Route, r.Handler)
	}
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
