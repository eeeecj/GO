package main

import (
	"context"
	"fmt"
	"sync"
)

func GenerateNatural2(ctx context.Context, wg *sync.WaitGroup) chan int {
	ch := make(chan int)
	go func() {
		// 确定ctx被取消时，能够取消wg
		defer wg.Done()
		// 确保ctx被取消时能够关闭通道
		defer close(ch)
		for i := 2; ; i++ {
			select {
			// 输出i
			case ch <- i:
				// 取消ctx时返回
			case <-ctx.Done():
				return
			}
		}
	}()
	return ch
}

func filter(ctx context.Context, wg *sync.WaitGroup, in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		// 确保ctx被取消时，wg也被取消
		defer wg.Done()
		// 确保无论是输入管道被关闭，还是 ctx 被取消，只要素数筛退出，都会关闭输出管道。
		defer close(out)

		// 保证了输入管道被关闭时，循环能退出，不会出现死循环；
		for i := range in {
			if i%prime != 0 {
				select {
				case <-ctx.Done():
					return
				case out <- i:
				}
			}
		}
	}()
	return out
}

func main() {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	ch := GenerateNatural2(ctx, wg)
	// 输出100个素数
	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Printf("%v:%v\n", i+1, prime)
		wg.Add(1)
		ch = filter(ctx, wg, ch, prime)
	}
	cancel()
	wg.Wait()
}
