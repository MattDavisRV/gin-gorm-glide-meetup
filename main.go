package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB

type User struct {
	ID        int
	FirstName string `gorm:"column:FirstName"`
	LastName  string
	Age       int
}

func init() {

	var err error

	db, err = gorm.Open("sqlite3", "./users.db")
	db.LogMode(true)

	if err != nil {
		panic("Couldn't create db")
	}

	if !db.HasTable(&User{}) {
		db.CreateTable(new(User))
	}

}

func main() {

	router := gin.New()

	router.Use(func(c *gin.Context) {
		c.Set("Random", "This Is my String")
	})

	router.GET("/test", func(c *gin.Context) {

		c.String(http.StatusOK, "Testing...")
	})

	router.GET("/hello", sayHello)
	router.GET("/hello/:name", sayHelloToUser)

	router.POST("/user", addUser)
	router.GET("/user", listAllUsers)
	router.GET("/user/:id", getSingleUser)

	router.Run(":9000")
}

func getSingleUser(c *gin.Context) {

	id := c.Param("id")
	user := User{}

	db.Find(&user, id)

	c.JSON(http.StatusOK, user)

}

func listAllUsers(c *gin.Context) {

	users := []User{}

	db.Find(&users, "FirstName = 'Matt2'")

	c.JSON(http.StatusOK, users)

}

func addUser(c *gin.Context) {

	user := User{}

	if c.BindJSON(&user) != nil {
		c.String(http.StatusInternalServerError, "Unable to parse JSON")
		return
	}

	db.Create(&user)

	c.String(http.StatusOK, "Inserted record")

}

func sayHelloToUser(c *gin.Context) {

	name := c.Param("name")

	c.String(http.StatusOK, fmt.Sprintf("Hello %s!", name))

}

func sayHello(c *gin.Context) {

	string, _ := c.Get("Random")

	fmt.Println(string)

	c.String(http.StatusOK, "Hello!")

}
