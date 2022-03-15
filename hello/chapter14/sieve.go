package main

import (
	"fmt"
	"time"
)

func main() {
	// ch := make(chan int)
	// go generate(ch)
	// for prime := range ch {
	// 	fmt.Print(prime, " ")
	// 	ch1 := make(chan int)
	// 	go filter(ch, ch1, prime)
	// 	ch = ch1
	// }
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()

	go func() {
		for item := range ch {
			fmt.Print(item, " ")
		}
	}()
	time.Sleep(1e9)
}

func generate(ch chan int) {
	for i := 2; i < 1000; i++ {
		ch <- i
	}
	close(ch)
}

func filter(in chan int, out chan int, prime int) {
	for i := range in {
		if i%prime != 0 {
			out <- i
		}
	}
}
