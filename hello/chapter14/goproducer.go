package main

import "fmt"

func main() {
	num := make(chan int)
	go newGen(0, 10, num)
	done := make(chan bool)
	go getNum(num, done)

	<-done
}

func newGen(start int, count int, c chan<- int) {
	for i := start; i < count; i++ {
		c <- start
		start += count
	}
	close(c)
}

func getNum(in <-chan int, done chan<- bool) {
	for num := range in {
		fmt.Println(num)
	}
	done <- true
}
