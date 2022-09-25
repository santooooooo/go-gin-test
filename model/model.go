package model

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

type User struct {
	ID   int    `gorm:"primary_key;not null"`
	Name string `gorm:"type:varchar(200);not null"`
}

func getGormConnect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		// heroku上だとenvが取得できないことによるエラーでサーバーが停止してしまうため、heroku上のみコメントアウト
		//log.Fatal("Error loading .env file")
	}

	DBMS := os.Getenv("DBMS")
	USER := os.Getenv("DBUSER")
	PASS := os.Getenv("DBPASS")
	PROTOCOL := os.Getenv("DBPROTOCOL")
	DBNAME := os.Getenv("DBNAME")
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME

	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		fmt.Println(DBMS)
		fmt.Println(CONNECT)
		fmt.Println("error occurs!")
		panic(err.Error())
	}

	db.Set("gorm:table_option", "ENGINE=InnoDB")
	db.LogMode(true)
	db.SingularTable(true)
	db.AutoMigrate(&User{})

	fmt.Println("db connected: ", &db)
	return db
}

func insertUser(registerUser *User) User {
	db := getGormConnect()

	db.Create(&registerUser)
	defer db.Close()

	userInfo := User{
		ID:   registerUser.ID,
		Name: registerUser.Name,
	}
	return userInfo
}

func findAllUser() string {
	db := getGormConnect()
	var users []User

	db.Order("ID asc").Find(&users)
	defer db.Close()
	usersInfo, _ := json.Marshal(users)
	return string(usersInfo)
}

func FindUsers(c *gin.Context) {
	//var user = User{
	//		Name: "test",
	//	}

	//insertUser(&user)

	resultUsers := findAllUser()

	//for i := range resultUsers {
	//		fmt.Printf("index: %d, ユーザーID: %d, ユーザー名: %s\n",
	//		i, resultUsers[i].ID, resultUsers[i].Name)
	//	}

	c.JSON(200, resultUsers)
	return
}

func InsertUser(c *gin.Context) {
	var user = User{}
	user.Name = c.PostForm("name")
	userInfo := insertUser(&user)
	c.JSON(200, userInfo)
	return
}
