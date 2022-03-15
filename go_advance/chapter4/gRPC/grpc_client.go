package main

import (
	"context"
	"fmt"
	HelloService "go_advance/chapter4/gRPC/helloService"
	"log"

	"google.golang.org/grpc"
)

func main() {
	// 通过grpc.Dial与gRPC服务建立链接
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// NewHelloService基于建立的连接返回HelloServiceClient对象
	// 该方法由proto生成
	// HelloServiceClient为一个接口，可以通过该接口执行远程方法
	client := HelloService.NewHelloServiceClient(conn)
	// 执行远程方法，第一个参数传入空白上下文
	reply, err := client.Hello(context.Background(), &HelloService.String{Value: "Hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
