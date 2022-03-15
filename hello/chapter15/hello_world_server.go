package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

const forms = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
<form action="#" method="post" name="bar">
<input type="text" value="" placeholder="尼玛" name="in"/>
<input type="submit" value="submit"/>
</form>
</body>
</html>
	`

type HandleFnc func(http.ResponseWriter, *http.Request)

func LogPanic(f HandleFnc) HandleFnc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Fatal("[%v] caught err %v", r.RemoteAddr, x)
			}
		}()
		f(w, r)
	}
}

func main() {

	http.HandleFunc("/hello/", LogPanic(HelloName))
	http.HandleFunc("/shounthello/", LogPanic(shounthello))
	http.HandleFunc("/test1/", LogPanic(SimpleServer))
	http.HandleFunc("/test2/", LogPanic(FormServer))
	err := http.ListenAndServe("localhost:5000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
func SimpleServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}

func FormServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch r.Method {
	case "GET":
		io.WriteString(w, forms)
	case "POST":
		r.ParseForm()
		io.WriteString(w, r.Form["in"][0])
	}
}
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World"+r.URL.Path[1:])
}
func HelloName(w http.ResponseWriter, r *http.Request) {
	length := len("/hello/")
	fmt.Fprint(w, "Hello "+r.URL.Path[length:])
}

func shounthello(w http.ResponseWriter, r *http.Request) {
	length := len("/shounthello/")
	fmt.Fprint(w, "Hello "+strings.ToUpper(r.URL.Path[length:]))

}
