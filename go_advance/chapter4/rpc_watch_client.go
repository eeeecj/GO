package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("connection error")
	}

	go func() {
		var keychanged string
		// 注册远程监控程序，设定30s超时
		err := client.Call("KVStoreService.Watch", 30, &keychanged)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("watch:", keychanged)
	}()
	// 设定初始值，不触发watch
	err = client.Call("KVStoreService.Set", [2]string{"abc", "abc_v"}, new(struct{}))
	// 更改值，触发watch
	err = client.Call("KVStoreService.Set", [2]string{"abc", "abc_v1"}, new(struct{}))
	if err != nil {
		log.Fatal(err)
	}
	// 等待协程运行
	time.Sleep(time.Second * 3)
}
