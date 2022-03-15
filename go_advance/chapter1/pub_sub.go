package main

import (
	"fmt"
	"go_advance/chapter1/pubsub"
	"strings"
	"time"
)

func main() {
	p := pubsub.NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	all := p.Subscribe()
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	p.Publish("hello world")
	p.Publish("Hello golang")

	go func() {
		for msg := range all {
			fmt.Println("all", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println("golang", msg)
		}
	}()
	time.Sleep(3 * time.Second)
}
