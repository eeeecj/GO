package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	bks := make([]Book, 0)
	file, err := os.Open("./products.txt")
	if err != nil {
		log.Fatalf("Error %s opening file products.txt: ", err)
	}
	defer file.Close()
	var re string
	fmt.Println(fmt.Fscanln(file, &re))
	fmt.Println(re)
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		line = string(line[:len(line)-2])
		strs := strings.Split(line, ";")
		book := new(Book)
		book.title = strs[0]
		book.price, err = strconv.ParseFloat(strs[1], 32)
		if err != nil {
			fmt.Printf("Error in file: %v", err)
		}

		book.quantity, err = strconv.Atoi(strs[2])
		if err != nil {
			fmt.Printf("Error in file: %v", err)
		}
		bks = append(bks, *book)

	}

	fmt.Println("We have read the following books from the file: ")
	for _, bk := range bks {
		fmt.Println(bk)
	}
}

type Book struct {
	title    string
	price    float64
	quantity int
}
