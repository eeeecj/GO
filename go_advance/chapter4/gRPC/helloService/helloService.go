package HelloService

import "context"

type HelloServiceImp struct{}

// 构造函数需要满足proto生成的HelloServiceServer接口
func (p *HelloServiceImp) Hello(ctx context.Context, args *String) (*String, error) {
	reply := &String{Value: "Hello" + args.GetValue()}
	return reply, nil
}
