package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

func main() {
	// 建立tcp远程监听服务
	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
		log.Fatal("Listen TCP error")
	}
	// 建立缓存通道
	clientchan := make(chan *rpc.Client, 10)

	go func() {
		for {
			// 监听tcp
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal("Accept error")
			}
			// 将tcp转化为rpc接口
			clientchan <- rpc.NewClient(conn)
		}
	}()
	// 调用远程服务
	dosomework(clientchan)

}

func dosomework(clientchan chan *rpc.Client) {
	client := <-clientchan
	defer client.Close()
	var reply string
	err := client.Call("HelloService.Hello", "Hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
