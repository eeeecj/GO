package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"

	"go_advance/chapter4/safety/Token"
	HelloService "go_advance/chapter4/safety/helloservice"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	cred, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to add cert")
	}
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cred},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})
	server := grpc.NewServer(grpc.Creds(creds))

	// 创建验证对象
	auth := &Token.Authentication{User: "gopher", Password: "password"}
	Hello := &HelloService.HelloServiceImp{Auth: auth}
	HelloService.RegisterHelloServiceServer(server, Hello)

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	server.Serve(lis)
}
