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
	ID    int    `gorm:"primary_key;not null"`
	Name  string `gorm:"type:varchar(200);not null;unique"`
	Point int    `gorm: "type:int;not null;"`
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

func insertUser(registerUser *User) string {
	db := getGormConnect()

	result := db.Where("name = ?", registerUser.Name).First(&User{})
	if result.RowsAffected != 0 {
		return "このユーザーは既に存在しています"
	}

	db.Create(&registerUser)
	defer db.Close()

	jsonEncode, _ := json.Marshal(registerUser)
	return string(jsonEncode)
}

func InsertUser(c *gin.Context) {
	var user = User{}
	var existedUser string = "このユーザーは既に存在しています"
	user.Name = c.PostForm("name")
	userInfo := insertUser(&user)
	if userInfo == existedUser {
		c.JSON(403, userInfo)
		return
	}
	c.JSON(200, userInfo)
	return
}

func login(registerUser *User) string {
	var user = User{}
	db := getGormConnect()

	result := db.Where("name = ?", registerUser.Name).First(&user)
	if result.RowsAffected != 0 {
		jsonEncode, _ := json.Marshal(user)
		return string(jsonEncode)
	}

	return "正しいユーザー名を入力してください"
}

func Login(c *gin.Context) {
	var user = User{}
	var notUser string = "正しいユーザー名を入力してください"
	user.Name = c.PostForm("name")
	userInfo := login(&user)
	if userInfo == notUser {
		c.JSON(404, userInfo)
		return
	}
	c.JSON(200, userInfo)
	return
}

// pointの増加
func pointIncrement(registerUser *User) string {
	var user = User{}
	db := getGormConnect()
	var notUser string = "正しいユーザー名を入力してください"

	result := db.Where("name = ?", registerUser.Name).First(&user)
	if result.RowsAffected == 0 {
		return notUser
	}

	incrementPoint := user.Point + 1
	db.Model(&user).Update("point", incrementPoint)

	jsonEncode, _ := json.Marshal(user)
	return string(jsonEncode)
}

func PointIncrement(c *gin.Context) {
	var user = User{}
	var notUser string = "正しいユーザー名を入力してください"
	user.Name = c.PostForm("name")
	userInfo := pointIncrement(&user)
	if userInfo == notUser {
		c.JSON(404, userInfo)
		return
	}
	c.JSON(200, userInfo)
	return
}
