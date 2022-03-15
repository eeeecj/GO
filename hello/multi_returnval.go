package main

import "fmt"

func main() {
	y1, y2, y3 := func1(2, 3)
	fmt.Println(y1, y2, y3)
	fmt.Println(func2(2, 3))
}

func func1(x1 int, x2 int) (y1 int, y2 int, y3 int) {
	y1 = x1 + x2
	y2 = x1 * x2
	y3 = x1 - x2

	return
}

func func2(x1 int, x2 int) (int, int, int) {
	return x1 + x2, x1 * x2, x1 - x2
}
