package main

import (
	"fmt"
	"log"
	"net/http"
)

const AddForm = `
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
`

var store = NewUrlStore()

func main() {
	http.HandleFunc("/", Redict)
	http.HandleFunc("/add", Add)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func Redict(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	url := store.Get(key)
	if url == "" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}
func Add(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, AddForm)
		return
	}
	key := store.Put(url)
	fmt.Fprintf(w, "http://localhost:5000/%s", key)
}
