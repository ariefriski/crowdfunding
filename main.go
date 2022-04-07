package main

import (
	"crowdfunding/handler"
	"crowdfunding/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var a int = 0
// var point *int = &a
// fmt.Print(point)

func main() {
	dsn := "root:123123@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())

	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("/api/v1/")

	api.POST("/users", userHandler.RegisterUser)

	router.Run()

	// userInput := user.RegisterUserInput{}
	// userInput.Name = "Nadya Apriliani"
	// userInput.Email = "Nadya@gmail.com"
	// userInput.Occupation = "Mahasiswi"
	// userInput.Password = "123123"

	// userService.RegisterUser(userInput)

}