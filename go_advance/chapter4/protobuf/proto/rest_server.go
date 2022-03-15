package main

import (
	"context"
	"go_advance/chapter4/protobuf/proto/rest"
	"log"
	"net/http"

	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	err := rest.RegisterRestServiceHandlerFromEndpoint(ctx, mux, "localhost:1234", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(":8000", mux)
}
