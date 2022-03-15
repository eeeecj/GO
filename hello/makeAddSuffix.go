package main

import (
	"fmt"
	"strings"
)

func main() {
	jpg_s := makeAddSuffix(".jpg")
	fmt.Println(jpg_s("files"))
}

func makeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
