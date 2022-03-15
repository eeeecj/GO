package main

import (
	"fmt"
	"hello/chapter14/buildlazyevaluator"
)

func main() {
	// evenFunc := func(state buildlazyevaluator.Any) (buildlazyevaluator.Any, buildlazyevaluator.Any) {
	// 	os := state.(int)
	// 	ns := os + 2
	// 	return os, ns
	// }
	// even := buildlazyevaluator.BuildIntLazyEvaluator(evenFunc, 4)
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("%vth even:%v\n", i, even())
	// }

	fibolaci := func(state buildlazyevaluator.Any) (buildlazyevaluator.Any, buildlazyevaluator.Any) {
		os := state.([]uint64)
		v1 := os[0]
		v2 := os[1]
		ns := []uint64{v2, v1 + v2}
		return v1, ns
	}
	fib := BuildUint64lazyevaluator(fibolaci, []uint64{0, 1})
	for i := 0; i < 25; i++ {
		fmt.Printf("Fib nr %v: %v\n", i, fib())
	}
}

func BuildUint64lazyevaluator(evenFunc buildlazyevaluator.EvalFunc, state buildlazyevaluator.Any) func() uint64 {
	ef := buildlazyevaluator.BuildLazyEvaluator(evenFunc, state)
	return func() uint64 {
		return ef().(uint64)
	}
}
