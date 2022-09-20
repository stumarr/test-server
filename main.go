package main

import (
	"github.com/gorilla/mux"
	"net/http"
	h "southwinds.dev/http"
)

func main() {
	// create a new server
	server := h.New("test", "1.0")
	// configure http handlers
	server.Http = func(router *mux.Router) {
		// add basic authentication
		router.Use(server.AuthenticationMiddleware)
		// add handler
		router.HandleFunc("/", doSomething).Methods("GET")
	}
	server.Serve()
}

func doSomething(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
}
