package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"hello/chapter12/stack"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	cal := new(stack.Stack)
	fmt.Println("Give a number, an operator (+, -, *, /), or q to stop:")
	for {
		input, err := buf.ReadString('\n')
		if err != nil {
			fmt.Println("Input error !")
			os.Exit(1)
		}
		t := input[:len(input)-2]
		switch {
		case t == "q":
			fmt.Println("Calculator stopped")
			return
		case t >= "0" && t <= "999999":
			i, _ := strconv.Atoi(t)
			cal.Push(i)
		case t == "+":
			q := cal.Pop()
			p := cal.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, t, q, p+q)
		case t == "-":
			q := cal.Pop()
			p := cal.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, t, q, p-q)

		case t == "*":
			q := cal.Pop()
			p := cal.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, t, q, p*q)

		case t == "/":
			q := cal.Pop()
			p := cal.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, t, q, p/q)
		default:
			fmt.Println("No valid input")
		}
	}
}
