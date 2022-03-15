package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("starting server....")

	listener, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		go doServer(conn)
	}
}

func doServer(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("读取的字符串为:", string(buf[:len]))
	}
}
