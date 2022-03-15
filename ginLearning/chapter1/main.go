package main

import (
	"ginLearning/chapter1/controllers/hello"
	"ginLearning/chapter1/controllers/pings"
	"ginLearning/chapter1/routers"
)

func main() {
	routers.Include("/h", hello.Router)
	routers.Include("/p", pings.Router)
	r := routers.Init()
	r.Run(":8080")
}
