package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 使用
func worker(ctx context.Context, wg *sync.WaitGroup) error {
	// 保证wg关闭
	defer wg.Done()
	for {
		select {
		default:
			fmt.Println("hello world")
			// 保证ctx上下文关闭，Done（）返回关闭通道
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func main() {
	// 定义超时时间的上下文，context.Background()返回没有关闭时间的空上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// 定义组
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		// 添加组任务
		wg.Add(1)
		go worker(ctx, wg)
	}
	time.Sleep(time.Second)
	// 执行cancel关闭ctx上下文，并关闭所有协程中的子上下文
	cancel()
	// 等待结束时间
	wg.Wait()
}
