package HelloWordSafe

import "net/rpc"

//  RPC 服务的接口规范分为三个部分：
// 首先是服务的名字，
// 然后是服务要实现的详细方法列表，
// 最后是注册该类型服务的函数。

// 简化服务端
const HelloServiceName = "HelloWordSafe/pkg.HelloService"

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

func RegisterHelloServer(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

// 定义客户端结构
type HelloService struct{}

// 定义服务端处理函数
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

// 简化客户端
// 新增客户端类型
type HelloServiceClient struct {
	Client *rpc.Client
}

// 规定客户端类型满足HelloServiceInterface接口
// 客户端用户就可以直接通过接口对应的方法调用 RPC 函数
var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DailHelloService(network, addr string) (*HelloServiceClient, error) {
	// 拨号访问
	c, err := rpc.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	// 返回客户端连接
	return &HelloServiceClient{Client: c}, nil
}

// 定义RPC函数，实现相应的接口
func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}
