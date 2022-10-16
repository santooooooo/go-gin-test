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

	// ユーザー全体の情報の閲覧
	router.GET("/users", model.FindUsers)

	// ユーザーの登録
	router.POST("/user", model.InsertUser)

	// ユーザーのログイン
	router.POST("/login", model.Login)

	// ユーザー全体の情報の閲覧
	router.POST("/point/increment", model.PointIncrement)

	router.Run(":" + os.Getenv("PORT"))
}
