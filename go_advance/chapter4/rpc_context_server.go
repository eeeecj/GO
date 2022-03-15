package main

import (
	helloworldsafecontext "go_advance/chapter4/helloworldSafeContext"
	"log"
	"net"
	"net/rpc"
)

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// 使用协程，为每一个连接建立一个协程
		go func() {
			defer conn.Close()
			p := rpc.NewServer()
			// 注册服务
			p.RegisterName("helloworldsafecontext/HelloService", &helloworldsafecontext.HelloService{Conn: conn})
			p.ServeConn(conn)
		}()
	}
}
