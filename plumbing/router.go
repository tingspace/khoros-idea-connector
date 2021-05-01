package plumbing

import (
	"fmt"
	"net/http"
)

type RouteItem struct {
	Route   string
	Handler func(http.ResponseWriter, *http.Request)
}

func Routes() []RouteItem {
	return []RouteItem{
		{"/health", healthHandler},
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request from: ", r.Host)
	w.WriteHeader(http.StatusOK)
}
