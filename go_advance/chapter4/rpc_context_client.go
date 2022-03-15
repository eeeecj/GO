package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	// 建立远程连接
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("dailing error")
	}
	var reply string
	// 远程登录
	err = client.Call("helloworldsafecontext/HelloService.Login", "user:Password", &reply)
	if err != nil {
		log.Fatal(err)
	}
	// 执行远程方法
	err = client.Call("helloworldsafecontext/HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
