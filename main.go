package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello Gin!")
	})
	router.GET("/users", model.findAllUser)

	router.Run(":8080")
}
