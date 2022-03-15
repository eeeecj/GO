package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var Options = make(map[string]func(c *gin.RouterGroup))

func Include(s string, f func(c *gin.RouterGroup)) {
	fmt.Println(s)
	Options[s] = f
}
func Init() *gin.Engine {
	r := gin.Default()
	fmt.Println(Options)
	for k, v := range Options {
		group := r.Group(k)
		v(group)
	}
	return r
}
