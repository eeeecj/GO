package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"go_advance/chapter4/safety/Token"
	HelloService "go_advance/chapter4/safety/helloservice"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// 读取证书
	cred, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		log.Fatal(err)
	}
	// 创建证书池
	certPool := x509.NewCertPool()

	ca, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		log.Fatal(err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("Failed to parse certificate")
	}

	// 创建证书容器
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cred},
		ServerName:   "server.io",
		RootCAs:      certPool,
	})

	// 启用证书验证和token验证
	auth := &Token.Authentication{User: "gopher", Password: "password"}
	conn, err := grpc.Dial("localhost:1234", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(auth))

	if err != nil {
		log.Fatal(err)
	}
	client := HelloService.NewHelloServiceClient(conn)

	reply, err := client.Hello(context.Background(), &HelloService.String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)

}
