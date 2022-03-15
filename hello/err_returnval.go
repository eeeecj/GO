package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	y1, e := mysqrt(5)
	if e != nil {
		fmt.Print(e)
		return
	}
	fmt.Println(y1)
	y2, e := mysqrt2(-8)
	if e != nil {
		fmt.Print(e)
		return
	}
	fmt.Print(y2)
}

func mysqrt(x1 float64) (y1 float64, e error) {
	if x1 < 0 {
		y1 = -1
		e = errors.New("开放不能为负数")
		return
	}
	y1 = math.Sqrt(x1)
	e = nil
	return
}
func mysqrt2(x1 float64) (float64, error) {
	if x1 < 0 {
		return -1, errors.New("开方不能为负数")
	}
	return math.Sqrt(x1), nil
}
