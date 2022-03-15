package main

import (
	"context"
	"go_advance/chapter4/protobuf/proto/rest"
	"net"
	reflect "reflect"

	reflection "google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

type RestServiceImp struct{}

func (r *RestServiceImp) Get(ctx context.Context, message *rest.StringMessage) (*rest.StringMessage, error) {
	return &rest.StringMessage{Value: "Get hi:" + message.Value}, nil
}

func (r *RestServiceImp) Post(ctx context.Context, message *rest.StringMessage) (*rest.StringMessage, error) {
	return &rest.StringMessage{Value: "Post hi:" + message.Value}, nil
}
func main() {
	c("sss")
	grpcserver := grpc.NewServer()
	rest.RegisterRestServiceServer(grpcserver, new(RestServiceImp))
	lis, _ := net.Listen("tcp", ":1234")

	reflection.Register(grpcserver)
	grpcserver.Serve(lis)
}

func c(x interface{}) {
	xv := reflect.ValueOf(x)
	t := xv.NumField()
	println(t)
}
