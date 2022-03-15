package main

import (
	"context"
	PubSubStream "go_advance/chapter4/grpc/pubsub"
	"log"

	"google.golang.org/grpc"
)

func main() {
	// 建立grpc连接
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	// 使用proto生成方法构建客户端
	client := PubSubStream.NewPubSubServiceClient(conn)
	// 发送符合筛选条件的消息
	// 使用_忽略空返回值
	_, err = client.Publish(context.Background(), &PubSubStream.String{Value: "golang:hello go"})

	if err != nil {
		log.Fatal(err)
	}
	// 发送不符合筛选条件的消息
	_, err = client.Publish(context.Background(), &PubSubStream.String{Value: "docker:hello Docker"})
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Publish(context.Background(), &PubSubStream.String{Value: "golang:hello Docker"})
	if err != nil {
		log.Fatal(err)
	}
}
