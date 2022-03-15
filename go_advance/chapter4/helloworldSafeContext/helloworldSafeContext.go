package helloworldsafecontext

import (
	"fmt"
	"log"
	"net"
)

type HelloService struct {
	Conn    net.Conn
	isLogin bool
}

func (p *HelloService) Login(request string, reply *string) error {
	if request != "user:Password" {
		return fmt.Errorf("auth failed")
	}
	log.Println("login ok")
	p.isLogin = true

	return nil
}

func (p *HelloService) Hello(request string, reply *string) error {
	// 判断登录状态
	if !p.isLogin {
		return fmt.Errorf("please login")
	}
	*reply = "hello:" + request + ",from" + p.Conn.RemoteAddr().String()
	return nil
}
