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
	"net"

	gprc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
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
		log.Fatal("failed to parse certificate")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cred},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})
	server := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(gprc_middleware.ChainUnaryServer(filter)))

	auth := &Token.Authentication{User: "gopher", Password: "password"}

	Hello := &HelloService.HelloServiceImp{Auth: auth}

	HelloService.RegisterHelloServiceServer(server, Hello)

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	server.Serve(lis)
}

func filter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("filter", info)
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic:%v", r)
		}
	}()
	return handler(ctx, req)
}
