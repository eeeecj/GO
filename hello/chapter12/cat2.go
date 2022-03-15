package main

import (
	"flag"
	"fmt"
	"os"
)

var anumberFlag = flag.Bool("n", false, "list")
var i int = 0

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		Cat(os.Stdin)
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if f == nil {
			fmt.Fprintf(os.Stderr, "cannot:cannot reading from %s:%s\n", flag.Arg(i), err.Error())
		}
		Cat(f)
		f.Close()
	}
}

func Cat(f *os.File) {
	var buf [512]byte
	for {
		switch nr, err := f.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "%s", err.Error())
			os.Exit(1)
		case nr == 0:
			return
		case nr > 0:
			if *anumberFlag {
				fmt.Fprintf(os.Stdout, "%5d %s", i, buf[0:nr])
				i++
			} else {
				fmt.Fprintf(os.Stdout, "%s", buf[0:nr])
			}

		}
	}
}
