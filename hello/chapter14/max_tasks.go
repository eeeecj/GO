package main

import "fmt"

const MAXREQ = 50

var sem = make(chan int, MAXREQ)

type Request2 struct {
	a, b   int
	replyc chan int
}

func Process(r *Request2) {
	r.replyc <- r.a + r.b
}

func handle(r *Request2) {
	sem <- 1
	Process(r)
	<-sem
}

func server1(r chan *Request2) {
	for {
		req := <-r
		go handle(req)
	}
}
func main() {
	s := make(chan *Request2)
	go server1(s)
	for i := 0; ; i++ {
		req := &Request2{i, i + 100, make(chan int)}
		s <- req
		fmt.Println(<-req.replyc)
	}
}
