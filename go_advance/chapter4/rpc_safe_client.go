package main

import (
	"fmt"
	HelloWordSafe "go_advance/chapter4/HelloWorldSafe"
	"log"
)

func main() {
	// 获取访问客户端
	client, err := HelloWordSafe.DailHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Dailing err")
	}
	var reply string
	// 访问包中函数
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
