package main

import (
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jqs7/demo/controllers"
)

func main() {
	g := gin.Default()
	v1 := g.Group("/v1")
	v1.Use(cors.Default())
	v1.GET("/", controllers.Hello)
	g.Run(":8090")
}
