package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"strings"
)

type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() { // main函数，是程序执行的入口
	// a := 1
	str := "ksak dkskla ;dja"

	sl := strings.Fields(str)
	fmt.Printf("%v\n", sl)
	fmt.Println()
	fmt.Printf("%v\n", strings.Split(str, " "))
	// fmt.Println("Hello World!") // 在终端打印 Hello World!
	// fmt.Println("%s", runtime.Version())
	// fmt.Println(C.random())
	// fmt.Printf("%d", &a)

	fmt.Println(KB)

	for i := ""; i != "aaaa"; {
		println("this is first for")
		i += "a"
	}
}
