package main

import (
	"fmt"
	"time"

	"golang.org/x/net/websocket"
)

func main() {
	ws, err := websocket.Dial("ws://localhost:5000/websocket", "", "http://localhost/")
	if err != nil {
		panic("Dial: " + err.Error())
	}
	go readFormHandler(ws)
	time.Sleep(5e9)
	ws.Close()
}

func readFormHandler(ws *websocket.Conn) {
	buf := make([]byte, 100)
	for {
		if _, err := ws.Read(buf); err != nil {
			fmt.Printf("%s\n", err.Error())
			break
		}
	}
}
