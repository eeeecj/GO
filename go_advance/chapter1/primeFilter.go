package main

import (
	"fmt"
	"time"
)

// 数字产生器，返回一个通道
func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// 构造素数筛，删除能被素数整除的数，则输出通道
func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	// 使用并发创建素数整除，每一个素数将创建一个并发程序，所有的数将与in通道的素数整除
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	ch := GenerateNatural()
	start := time.Now()
	for i := 0; i <= 100; i++ {
		prime := <-ch
		fmt.Printf("%v:%v\n", i+1, prime)
		ch = PrimeFilter(ch, prime)
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}
