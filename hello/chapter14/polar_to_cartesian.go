package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Polar struct {
	radis float64
	o     float64
}

type Certificate struct {
	x float64
	y float64
}

func main() {
	question := make(chan Polar)
	defer close(question)

	answer := createCertificate(question)
	defer close(answer)
	interact(question, answer)
}

func createCertificate(polar chan Polar) chan Certificate {
	answer := make(chan Certificate)
	answer1 := new(Certificate)
	go func() {
		for {
			point := <-polar
			answer1.x = point.radis * math.Cos(point.o*math.Pi/180)
			answer1.y = point.radis * math.Sin(point.o*math.Pi/180)
			answer <- *answer1
		}
	}()
	return answer
}

func interact(question chan Polar, answer chan Certificate) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Radius and angle: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = line[0 : len(line)-1]
		if numbers := strings.Fields(line); len(numbers) == 2 {
			polars, err := stringTofloat(numbers)
			if err != nil {
				fmt.Fprintln(os.Stderr, "invalid number")
				continue
			}
			question <- Polar{polars[0], polars[1]}
			coord := <-answer
			fmt.Printf("r:%0.2f o:%0.2f x:%0.2f y:%0.2f", polars[0], polars[1], coord.x, coord.y)
			fmt.Println()
		} else {
			fmt.Fprintln(os.Stderr, "invalid input")
		}
	}
	fmt.Println()
}

func stringTofloat(fl []string) ([]float64, error) {
	var floats []float64
	for _, number := range fl {
		if x, err := strconv.ParseFloat(number, 64); err != nil {
			return nil, err
		} else {
			floats = append(floats, x)
		}
	}
	return floats, nil
}
