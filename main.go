package main

import (
	"main/model"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello Gin!")
	})
	router.GET("/users", model.FindUsers)

	router.POST("/users", model.InsertUser)
	//name := c.PostForm("name")
	//name := c.Query("name")
	//name := c.Param("name")

	router.Run(":80")
}
