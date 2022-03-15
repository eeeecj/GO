package main

import (
	"flag"
	"fmt"
	"net"
	"syscall"
)

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		fmt.Println("Usage:host:port")
	}
	hostAndport := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
	listener := initListen(hostAndport)
	for {
		conn, err := listener.Accept()
		checkErr(err, "Accept:")
		go connectionHandler(conn)
	}

}

func initListen(hostAndport string) *net.TCPListener {
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndport)
	checkErr(err, "Resolving host and port failed:"+hostAndport)
	listener, err := net.ListenTCP("tcp", serverAddr)
	checkErr(err, "Listen tcp:")
	println("Listening to :", listener.Addr().String())
	return listener
}

func checkErr(err error, s string) {
	if err != nil {
		panic("error:" + s + " " + err.Error())
	}
}

func connectionHandler(conn net.Conn) {
	connForm := conn.RemoteAddr().String()
	fmt.Println("Remote From:", connForm)
	for {
		var buf []byte = make([]byte, 26)
		length, err := conn.Read(buf[0 : len(buf)-2])
		buf[len(buf)-1] = 0
		switch err {
		case nil:
			handle(length, buf)
		case syscall.EAGAIN:
			continue
		default:
			goto CLOSE
		}
	}
CLOSE:
	err := conn.Close()
	println("close connection:", connForm)
	checkErr(err, "关闭服务器失败")
}

func handle(length int, buf []byte) {
	if length > 0 {
		fmt.Printf("<")
		for i := 0; i < length; i++ {
			if buf[i] != 0 {
				fmt.Printf("%s", string(buf[i]))
			}
		}
		fmt.Printf(">")
	}
}
