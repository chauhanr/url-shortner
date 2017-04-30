package main

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Handle(ADD_URL, http.HandlerFunc(Add))
	router.Handle(REDIRECT_URL, http.HandlerFunc(Retrieve))

	http.ListenAndServe(SERVER_PORT, context.ClearHandler(router))
}
