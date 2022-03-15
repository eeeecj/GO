package main

import (
	"hello/chapter15/rpc/rpc_objects"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

func main() {
	calc := new(rpc_objects.Args)
	rpc.Register(calc)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatal("Starting rpc-server error:", err.Error())
	}
	go http.Serve(listener, nil)
	time.Sleep(100e9)
}
