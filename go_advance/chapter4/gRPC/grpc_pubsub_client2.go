package main

import (
	"context"
	"fmt"
	PubSubStream "go_advance/chapter4/grpc/pubsub"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	// 建立订阅客户端链接
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}
	client := PubSubStream.NewPubSubServiceClient(conn)
	// 获得获取信息的流
	stream, err := client.Subscribe(context.Background(), &PubSubStream.String{Value: "golang:"})

	if err != nil {
		log.Fatal(err)
	}

	for {
		// 等待接收信息
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue())
	}

}
