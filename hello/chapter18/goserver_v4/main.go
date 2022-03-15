package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var store = NewUrlStore("store.json")

func main() {
	http.HandleFunc("/", Redict)
	http.HandleFunc("/add", Add)
	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal("serving start error:", err.Error())
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
		w.Header().Set("content-type", "text/html; charset=utf-8")
		fmt.Fprint(w, AddForm)
		return
	}
	key := store.Put(url)
	http.Redirect(w, r, strings.Join([]string{"http://localhost:5000/", key}, ""), http.StatusFound)
}

const AddForm = `
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
`
