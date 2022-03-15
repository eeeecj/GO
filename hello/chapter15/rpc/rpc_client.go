package main

import (
	"fmt"
	"hello/chapter15/rpc/rpc_objects"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:5000")
	if err != nil {
		log.Fatal("Error daily:", err.Error())
	}

	args := &rpc_objects.Args{7, 8}

	var reply int
	err = client.Call("Args.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Args Error:", err.Error())
	}
	fmt.Printf("Args:%d+%d=%d", args.M, args.N, reply)
}
