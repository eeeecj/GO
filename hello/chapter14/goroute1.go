package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start main")
	go longwait()
	go shortwait()
	fmt.Println("start sleep in main")
	time.Sleep(10 * 1e9)
	fmt.Println("end sleep in main")
}

func longwait() {
	fmt.Println("begin long wait")
	time.Sleep(5 * 1e9)
	fmt.Println("end long wait")
}

func shortwait() {
	fmt.Println("begin short wait")
	time.Sleep(2 * 1e9)
	fmt.Println("end short wait")
}
