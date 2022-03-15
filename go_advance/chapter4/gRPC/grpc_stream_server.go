package main

import (
	helloServiceStream "go_advance/chapter4/gRPC/helloServicestream"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	grpcserver := grpc.NewServer()
	helloServiceStream.RegisterHelloServiceStramServer(grpcserver, new(helloServiceStream.HelloServiceStramImp))
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcserver.Serve(lis)
}
