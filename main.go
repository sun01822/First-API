package main

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


type User struct{
	gorm.Model
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
	Address string `json:"address"`
}

var DB *gorm.DB 

func main(){
	e := echo.New()

	// Initiate DB 
	dsn := "root:@Sun2021Yes#@tcp(127.0.0.1:3306)/userlist?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	DB = d
	fmt.Println("DB connected")
	
	// Create and migrate the DB 
	DB.AutoMigrate(&User{})
	fmt.Println("DB migrated")
	
	// Routes
	e.GET("/users", GetAllUsers)
	e.POST("/users", CreateUser)
	e.DELETE("/users/:id", DeleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}


func CreateUser(c echo.Context) error{
	user := &User{}
	if err := c.Bind(user); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := DB.Create(&user).Error; err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, "User created successfully")
}


func GetAllUsers(c echo.Context) error{
	var users []User
	if err := DB.Find(&users).Error; err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}


func DeleteUser(c echo.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))
	var user User
	if err := DB.First(&user, id).Error; err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if err := DB.Delete(&user).Error; err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "User deleted successfully")
}


