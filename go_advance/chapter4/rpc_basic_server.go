package main

import (
	helloword "go_advance/chapter4/hello_word"
	"log"
	"net"
	"net/rpc"
)

func main() {
	// 注册名称为HelloService的服务，将对象中的方法注册为RPC函数
	// 所有的注册方法均会放在“HelloService”服务空间下
	rpc.RegisterName("HelloService", new(helloword.HelloServiece))
	// 使用TCP方法建立连接
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen to failed")
		return
	}
	// 接受请求
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error")
	}
	// 通过TCP建立连接
	rpc.ServeConn(conn)
}
