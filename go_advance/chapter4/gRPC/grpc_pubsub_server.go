package main

import (
	PubSubStream "go_advance/chapter4/grpc/pubsub"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// 建立grpc链接
	grpcserver := grpc.NewServer()
	// 注册服务，注意此处需要用构建函数作为第二个输入值，否则会出现空指针
	PubSubStream.RegisterPubSubServiceServer(grpcserver, PubSubStream.NewPubSubService())
	lis, err := net.Listen("tcp", ":1234")

	if err != nil {
		log.Fatal("Listen error")
	}
	grpcserver.Serve(lis)
}
