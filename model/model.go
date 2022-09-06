package model

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID   int    `gorm:"primary_key;not null"`
	Name string `gorm:"type:varchar(200);not null"`
}

func getGormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "user"
	PASS := "gogin"
	PROTOCOL := "tcp(db)"
	DBNAME := "go-gin"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
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

func insertUser(registerUser *User) {
	db := getGormConnect()

	db.Create(&registerUser)
	defer db.Close()
}

func findAllUser() []User {
	db := getGormConnect()
	var users []User

	db.Order("ID asc").Find(&users)
	defer db.Close()
	return users
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
	insertUser(&user)
	c.JSON(200, "Success")
	return
}
