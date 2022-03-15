package main

import (
	rpckv "go_advance/chapter4/rpcKV"
	"log"
	"net"
	"net/rpc"
)

func main() {
	rpc.RegisterName("KVStoreService", rpckv.NewKVStoreService())

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error")
		}
		go rpc.ServeConn(conn)
	}
}
