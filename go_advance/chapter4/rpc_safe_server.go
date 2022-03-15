package main

import (
	HelloWorldSafe "go_advance/chapter4/HelloWorldSafe"
	"log"
	"net"
	"net/rpc"
)

func main() {
	// 使用包中函数注册RPC服务
	HelloWorldSafe.RegisterHelloServer(new(HelloWorldSafe.HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTcp err")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error")
		}
		go rpc.ServeConn(conn)
	}
}
