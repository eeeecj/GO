package main

import (
	HelloService "go_advance/chapter4/gRPC/helloService"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// 通过grpc.NewServer()构造服务对象
	grpcServer := grpc.NewServer()
	// 使用proto生成的注册函数注册我们想要执行的代码
	HelloService.RegisterHelloServiceServer(grpcServer, new(HelloService.HelloServiceImp))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	// 监听服务
	grpcServer.Serve(lis)
}
