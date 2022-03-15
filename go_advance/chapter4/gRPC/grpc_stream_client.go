package main

import (
	"context"
	"fmt"
	HelloServiceStream "go_advance/chapter4/gRPC/helloServicestream"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := HelloServiceStream.NewHelloServiceStramClient(conn)
	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			// 双向流发送数据
			if err := stream.Send(&HelloServiceStream.String{Value: "hi"}); err != nil && err != io.EOF {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		// 双向流接收数据
		reply, err := stream.Recv()
		if err != nil && err != io.EOF {
			log.Fatal("ee", err)
		}
		fmt.Println(reply.GetValue())
	}

}
