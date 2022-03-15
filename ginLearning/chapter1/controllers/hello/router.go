package hello

import "github.com/gin-gonic/gin"

func Router(g *gin.RouterGroup) {
	g.GET("/hello", Hello)
}
