package main

import (
	HelloWordSafe "go_advance/chapter4/HelloWorldSafe"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 建立连接
	HelloWordSafe.RegisterHelloServer(new(HelloWordSafe.HelloService))
	// 注册路由与处理函数
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		// 构造io.ReadWriteCloser接口
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	// 监听服务
	http.ListenAndServe(":1234", nil)
}
