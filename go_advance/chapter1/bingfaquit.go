package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cancel := make(chan bool)
	// 使用wg记录协程正常运行和退出
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go wait(wg, cancel)
	}
	// 睡眠一秒钟
	time.Sleep(time.Second)
	// 关闭通道
	close(cancel)
	// 等待任务完成，确保垃圾清理所有协程
	wg.Wait()
}

func wait(wg *sync.WaitGroup, cancel chan bool) {
	defer wg.Done()
	for {
		// 使用select保证协程退出
		select {
		// 正常输出
		default:
			fmt.Println("hello world")
			// 异常时退出程序
		case <-cancel:
			return
		}
	}
}
