package main

import (
	"main/model"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	env := godotenv.Load()
	if env != nil {
		// heroku上だとenvが取得できないことによるエラーでサーバーが停止してしまうため、heroku上のみコメントアウトしたよ
		//log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello Gin!")
	})

	router.GET("/users", model.FindUsers)

	router.POST("/users", model.InsertUser)

	router.Run(":" + os.Getenv("PORT"))
}
