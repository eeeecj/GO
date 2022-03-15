package main

func main() {
	c := make(chan int)
	go sum(12, 13, c)
	println(<-c)
}

func sum(x int, y int, z chan int) {
	z <- x + y
}
