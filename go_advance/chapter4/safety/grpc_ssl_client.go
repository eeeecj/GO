package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	HelloService "go_advance/chapter4/gRPC/helloService"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// 导入密钥
	certificate, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()

	ca, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal(err)
	}
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		// 此处跟openssl里面设置一样
		ServerName: "server.io",
		RootCAs:    certPool,
	})

	conn, err := grpc.Dial("localhost:1234", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}

	client := HelloService.NewHelloServiceClient(conn)

	resp, err := client.Hello(context.Background(), &HelloService.String{Value: "hello"})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
