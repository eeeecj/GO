package main

import (
	"bufio"
	"fmt"
	"os"
)

var nrchars, nrwords, nrlines int = 0, 0, 0

func main() {

	inputReader := bufio.NewReader(os.Stdin)
	println("Please enter some input, type S to stop: ")
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			println(err)
			return
		}
		if input == "S\r\n" {
			fmt.Println("Here are the counts:")
			fmt.Printf("Number of characters: %d\n", nrchars)
			fmt.Printf("Number of words: %d\n", nrwords)
			fmt.Printf("Number of lines: %d\n", nrlines)
			os.Exit(0)
		}
		Count(input)
	}
}

func Count(s string) {
	nrchars += len(s) - 2
	nrwords += len(s)
	nrlines++
}
