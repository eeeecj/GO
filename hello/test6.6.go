package main

import (
	"fmt"
)

func main() {
	printn(5)
	println()
	fmt.Printf("%d \n", jiecheng(8))

}

func printn(x int) {
	if x > 0 {
		fmt.Printf("%d ", x)
		x -= 1
		printn(x)
	}
	return

}

func jiecheng(x int) int {
	if x > 0 {
		return x * jiecheng(x-1)
	}
	return 1
}
