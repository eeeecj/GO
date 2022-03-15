package main

import (
	HelloWordSafe "go_advance/chapter4/HelloWorldSafe"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 使用封装的包建立连接
	HelloWordSafe.RegisterHelloServer(new(HelloWordSafe.HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen err")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept err")
		}
		// 使用ServeCodec代替ServerConn
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
