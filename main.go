package main

import (
	"crowdfunding/auth"
	"crowdfunding/handler"
	"crowdfunding/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var a int = 0
// var point *int = &a
// fmt.Print(*point)

func main() {
	dsn := "root:123123@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())

	}
	//Repo
	userRepository := user.NewRepository(db)
	authService := auth.NewService()
	
	//Service
	userService := user.NewService(userRepository)

	//Handler
	userHandler := handler.NewUserHandler(userService,authService)

	router := gin.Default()

	api := router.Group("/api/v1/")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/check_email", userHandler.CheckEmailAvailibility)
	api.POST("/avatars",userHandler.UploadAvatar)
	router.Run()

	// userInput := user.RegisterUserInput{}
	// userInput.Name = "Nadya Apriliani"
	// userInput.Email = "Nadya@gmail.com"
	// userInput.Occupation = "Mahasiswi"
	// userInput.Password = "123123"

	// userService.RegisterUser(userInput)

}
