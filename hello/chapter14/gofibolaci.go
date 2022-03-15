package main

import (
	"fmt"
	"time"
)

//----------方法一---------------
// func main() {
// 	var i int = 0
// 	c := make(chan int)
// 	term := 25
// 	start := time.Now()
// 	go fibiterms(term, c)
// 	for {
// 		if res, ok := <-c; ok {
// 			fmt.Println(res)
// 			i++
// 		} else {
// 			end := time.Now()
// 			delta := end.Sub(start)
// 			fmt.Printf("运行时间%s\n", delta)
// 			os.Exit(0)
// 		}
// 	}
// }

// func fibiterms(term int, c chan int) {
// 	for i := 0; i < term; i++ {
// 		c <- fibilaci(i)
// 	}
// 	close(c)
// }

// func fibilaci(i int) (res int) {
// 	if i <= 1 {
// 		res = 1
// 	} else {
// 		res = fibilaci(i-1) + fibilaci(i-2)
// 	}
// 	return
// }
//--------------方法二-----------

// func main() {
// 	n := 25
// 	c := make(chan int)
// 	start := time.Now()
// 	go fibilaci(n, c)
// 	for i := range c {
// 		fmt.Println(i)
// 	}
// 	end := time.Now()
// 	delta := end.Sub(start)
// 	fmt.Printf("运行时间为%s\n", delta)
// }

// func fibilaci(n int, c chan int) {
// 	x, y := 1, 1
// 	for i := 0; i < n; i++ {
// 		c <- x
// 		x, y = y, x+y
// 	}
// 	close(c)
// }

//-----------方案三------------

func fib() <-chan int {
	x := make(chan int, 2)
	a, b, c := func() (chan int, chan int, chan int) {
		a, b, c := make(chan int, 2), make(chan int, 2), make(chan int, 2)
		go func() {
			for {
				t := <-x
				a <- t
				b <- t
				c <- t
			}
		}()
		return a, b, c
	}()
	go func() {
		x <- 0
		x <- 1
		<-a
		for {
			x <- <-a + <-b
		}
	}()
	<-c
	return c
}

func main() {
	start := time.Now()
	c := fib()
	n := 25
	for i := 0; i < n; i++ {
		fmt.Println(<-c)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("计算时间%s\n", delta)
}
