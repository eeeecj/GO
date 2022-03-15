package pubsub

import (
	"sync"
	"time"
)

type (
	subscriber chan interface{}         //定义管道
	topicFunc  func(v interface{}) bool // 定义过滤器
)

type Publisher struct {
	m           sync.RWMutex             //读写锁
	buffer      int                      //缓存
	timeout     time.Duration            //有效时间
	subscribers map[subscriber]topicFunc //订阅者信息
}

// 运行逻辑：
// 1、构建发布者，存储订阅者信息
// 2、构建一个订阅者，此时生成一个通道，该通道为键值，筛选函数作为值
// 3、发布一条信息，此时需要读取所有的订阅者信息
// 4、向订阅者通道发送信息、执行筛选函数等

// 构建发布者对象，可设置发布超时时间和缓存数量
func NewPublisher(pubulishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     pubulishTimeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

// 添加一个订阅者，订阅全部主题
func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

// 添加一个订阅者，订阅筛选后的主题
func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	// 使用通道作为key，用于储存筛选函数
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

// 退出订阅
func (p *Publisher) Evict(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.subscribers, sub)
	close(sub)
}

// 发布主题
func (p *Publisher) Publish(v interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.SendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

// 发送信息，可容忍一定时间超时
func (p *Publisher) SendTopic(sub chan interface{}, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	// 执行筛选函数，若函数执行失败则返回
	if topic != nil && !topic(v) {
		return
	}
	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}

}

// 关闭发布者通道，并关闭所有订阅者通道
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}
