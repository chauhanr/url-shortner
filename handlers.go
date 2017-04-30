package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var store = NewURLStore("url.json")

const AddForm = `
<html>
<body>
<form method="POST" action="/url/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
</body>
</html>
`

// Add func in handlers will handle the saving of the url to the store.
func Add(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		fmt.Fprint(w, AddForm)
		return
	}
	key := store.Put(url)
	fmt.Fprintf(w, "http://localhost%s/%s", SERVER_PORT, key)
}

// Retrieve function will retrieve the full url and then forward the request to the actual url.
func Retrieve(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	url := store.Get(key)
	log.Printf("The URL to redirect : %s\n", url)

	if url == "" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}
