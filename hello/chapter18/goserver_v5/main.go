package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"strings"
)

var store Store

var rpcflag = flag.Bool("rpc", false, "enable RPC server")
var masterAddr = flag.String("master", "", "RPC master addr")
var host = flag.String("host", ":5000", "host addr")

func main() {
	flag.Parse()
	if *masterAddr != "" {
		store = NewProxyStore(*masterAddr)
	} else {
		store = NewUrlStore("store.json")
	}
	if *rpcflag {
		fmt.Println("sss")
		rpc.RegisterName("Store", store)
		fmt.Println("aaa")
		rpc.HandleHTTP()
	}

	http.HandleFunc("/", Redirect)
	http.HandleFunc("/add", Add)
	if err := http.ListenAndServe(*host, nil); err != nil {
		log.Fatal(err.Error())
	}
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	var url string
	if err := store.Get(&key, &url); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if url == "" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

func Add(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		w.Header().Set("Content-Type", "text/html;charset=utf-8")
		fmt.Fprint(w, AddForm)
		return
	}
	var key string
	if err := store.Put(&url, &key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, strings.Join([]string{"http://localhost", *host, "/", key}, ""), http.StatusFound)
}

const AddForm = `
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
`
