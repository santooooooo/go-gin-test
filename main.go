package main

import (
	"main/model"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	env := godotenv.Load()
	port := ""
	if env != nil {
		//log.Fatal("Error loading .env file")
		port = "80"
	} else {
		port = os.Getenv("PORT")
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello Gin!")
	})
	router.GET("/users", model.FindUsers)

	router.POST("/users", model.InsertUser)
	//name := c.PostForm("name")
	//name := c.Query("name")
	//name := c.Param("name")

	router.Run(":" + port)
}
