package main

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Handle(ADD_URL, http.HandlerFunc(Add))

	http.ListenAndServe(SERVER_PORT, context.ClearHandler(router))
}
