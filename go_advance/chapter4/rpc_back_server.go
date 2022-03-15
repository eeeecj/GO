package main

import (
	HelloWordSafe "go_advance/chapter4/HelloWorldSafe"
	"net"
	"net/rpc"
	"time"
)

func main() {
	// 注册服务
	rpc.RegisterName("HelloService", new(HelloWordSafe.HelloService))

	for {
		// 远程连接tcp服务
		conn, _ := net.Dial("tcp", "localhost:1234")
		if conn == nil {
			time.Sleep(time.Second)
			continue
		}
		// 监听rpc服务
		rpc.ServeConn(conn)
		conn.Close()
	}
}
