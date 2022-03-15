package main

import (
	"fmt"
	"os"
)

//------------方案一------------------
// func main() {
// 	var ok bool = true
// 	var i int
// 	c := make(chan int)
// 	go tel(c)
// 	for ok {
// 		if i, ok = <-c; ok {
// 			fmt.Println(i)
// 		}
// 	}
// }

// func tel(c chan int) {
// 	for i := 0; i < 5; i++ {
// 		c <- i
// 	}
// 	close(c)
// }

//----------------方案二----------------------

func main() {
	c := make(chan int)
	flag := make(chan bool)
	go tel(c, flag)
	for {
		select {
		case i := <-c:
			fmt.Println(i)
		case <-flag:
			os.Exit(0)
		}
	}
}

func tel(c chan int, flag chan bool) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	flag <- true
}
