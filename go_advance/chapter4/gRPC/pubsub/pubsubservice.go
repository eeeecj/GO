package PubSubStream

import (
	context "context"
	"go_advance/chapter1/pubsub"
	"strings"
	"time"
)

type PubSubService struct {
	pub *pubsub.Publisher
}

// 使用自定义的pubsub包
// 使用自定义的包的构建函数
func NewPubSubService() *PubSubService {
	return &PubSubService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

// 获取值并进行发送
func (p *PubSubService) Publish(ctx context.Context, arg *String) (*String, error) {
	p.pub.Publish(arg.GetValue())
	return &String{}, nil
}

// 获取筛选条件，并添加订阅者
func (p *PubSubService) Subscribe(arg *String, stream PubSubService_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})
	// 阻塞运行，等待从通道获取值，并将值发送至流中
	for v := range ch {
		if err := stream.Send(&String{Value: v.(string)}); err != nil {
			return err
		}
	}
	return nil
}
