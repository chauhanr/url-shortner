package main

import (
	"fmt"
	"net/http"
)

var store = NewURLStore()

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

// AddForm is the html snippet to show on screen.
