package HelloService

import (
	"context"

	"go_advance/chapter4/safety/Token"
)

type HelloServiceImp struct {
	Auth *Token.Authentication //注意大写
}

// 构造函数需要满足proto生成的HelloServiceServer接口
func (p *HelloServiceImp) Hello(ctx context.Context, args *String) (*String, error) {
	// 首行验证
	if err := p.Auth.Auth(ctx); err != nil {
		return nil, err
	}
	reply := &String{Value: "Hello" + args.GetValue()}
	return reply, nil
}
