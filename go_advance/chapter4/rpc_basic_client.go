package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	// 使用Dial远程拨号RPC服务
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dailing err")
	}
	var replay string
	// 调用RPC中的服务
	err = client.Call("HelloService.Hello", "hello", &replay)
	if err != nil {
		log.Fatal(err)
	}
	// replay获取到相应的结果
	fmt.Println(replay)
}
