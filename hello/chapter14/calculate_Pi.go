package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

//-------单核---------------
// func main() {
// 	start := time.Now()
// 	fmt.Println(calcPI(5000))
// 	end := time.Now()
// 	delta := end.Sub(start)
// 	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
// }

// func calcPI(n int) float64 {
// 	c := make(chan float64)
// 	for i := 0; i < n; i++ {
// 		go cal(c, float64(i))
// 	}
// 	res := 0.0
// 	for i := 0; i < n; i++ {
// 		res += <-c
// 	}
// 	return res
// }
// func cal(c chan float64, k float64) {
// 	c <- 4 * math.Pow(-1, k) / (2*k + 1)
// }

//--------多核---------------
const GNCU = 2

func main() {
	start := time.Now()
	runtime.GOMAXPROCS(2)
	fmt.Println(calcPI(5000))
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func calcPI(n int) float64 {
	c := make(chan float64)
	for i := 0; i < GNCU; i++ {
		go cal(c, i*n/GNCU, (i+1)*n/GNCU)
	}
	res := 0.0
	for i := 0; i < GNCU; i++ {
		res += <-c
	}
	return res
}
func cal(c chan float64, start int, end int) {
	res := 0.0
	for i := start; i < end; i++ {
		res += 4 * math.Pow(-1, float64(i)) / (2*float64(i) + 1)
	}
	c <- res
}
