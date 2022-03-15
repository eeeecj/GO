package main

import (
	"fmt"
	"time"
)

func main() {
	c := pump(5)
	go show(c)
	time.Sleep(1e9)
}

func pump(n int) chan int {
	c := make(chan int)
	go func() {
		for i := 0; i <= n; i++ {
			c <- i
		}
	}()
	return c
}

func show(c chan int) {
	for {
		fmt.Println(<-c)
	}
}
