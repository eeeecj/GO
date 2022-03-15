package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入你服务器名字")
	clientName, _ := inputReader.ReadString('\n')
	clientName = strings.Trim(clientName, "\r\n")
	for {
		fmt.Println("请输入你想传送的内容")
		input, _ := inputReader.ReadString('\n')
		trimInput := strings.Trim(input, "\r\n")
		if trimInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(clientName + ":" + trimInput))
	}
}
