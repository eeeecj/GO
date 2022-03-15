package HelloServiceStream

import (
	context "context"
	"io"
)

type HelloServiceStramImp struct{}

func (p *HelloServiceStramImp) Channel(stream HelloServiceStram_ChannelServer) error {
	for {
		// 服务端接收客户端数据
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		// 服务端对数据进行处理并返回客户端
		reply := &String{Value: "hello:" + args.GetValue()}
		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

func (p *HelloServiceStramImp) Hello(ctx context.Context, args *String) (*String, error) {
	reply := &String{Value: "Hello" + args.GetValue()}
	return reply, nil
}
