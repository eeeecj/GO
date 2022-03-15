package main

import "fmt"

type Request struct {
	a, b   int
	replyc chan int
}

func (r *Request) String() string {
	return fmt.Sprintf("%d+%d=%d", r.a, r.b, <-r.replyc)
}

type binOp func(a, b int) int

func run(Op binOp, req *Request) {
	req.replyc <- Op(req.a, req.b)
}

func server(Op binOp, server chan *Request, quit chan bool) {
	for {

		select {
		case req := <-server:
			go run(Op, req)
		case <-quit:
			return
		}
	}
}

func startserver(op binOp) (chan *Request, chan bool) {
	reqChan := make(chan *Request)
	quit := make(chan bool)
	go server(op, reqChan, quit)
	return reqChan, quit
}
func main() {
	adder, quit := startserver(func(a, b int) int { return a + b })
	const N = 10000
	var reqs [N]Request
	for i := 0; i < N; i++ {
		req := &reqs[i]
		req.a = i
		req.b = i + N
		req.replyc = make(chan int)
		adder <- req
	}
	for i := N - 1; i >= 0; i-- {
		if <-reqs[i].replyc != N+i*2 {
			fmt.Printf("失败在%v\n", i)
		} else {
			fmt.Println("Request ", i, "is ok!")
		}
	}
	// r1 := &Request{10, 1, make(chan int)}
	// r2 := &Request{12, 5, make(chan int)}
	// adder <- r1
	// adder <- r2
	// fmt.Println(r1, r2)

	quit <- true
	fmt.Printf("done")
}
