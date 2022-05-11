package handler

import (
	"crowdfunding/helper"
	"crowdfunding/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	//mengubah inputan JSON inputan menjadi Register User input
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Account failed registered", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadGateway, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input) //LOOOLl
	if err != nil {
		response := helper.APIResponse("Account failed registered", http.StatusUnprocessableEntity, "error", err.Error())
		c.JSON(http.StatusBadGateway, response)
		return
	}

	formatter := user.FormatUser(newUser, "tokencontoh")

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed ", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadGateway, response)
		return
	}

	loggedinUser, err := h.userService.LoginUser(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed ", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadGateway, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, "tokentokentoken")

	response := helper.APIResponse("Login Successfuly", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailibility(c *gin.Context) {
	//menerima inputan lalu mapping ke struct cek email
	//struct input  di passing ke service menghubungkannya ke repo
	//mengganti db
	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email Checking Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadGateway, response)
		return
	}

	IsEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "server error"}
		response := helper.APIResponse("Email Checking Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadGateway, response)
		return
	}

	data := gin.H{
		"is_available": IsEmailAvailable,
	}
	metaMessage := "email has been registered"

	if IsEmailAvailable {
		metaMessage = "Email is avalilabel"
	}
	response := helper.APIResponse(metaMessage, http.StatusOK, "Success", data)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) UploadAvatar(c *gin.Context){
	file,err := c.FormFile("avatar")
	if err != nil{
		data := gin.H{"is_uploaded":false}
		response := helper.APIResponse("Failed to upload avatar",http.StatusBadRequest,"error",data)
		
		c.JSON(http.StatusBadRequest, response)
		return
	}
	userID:= 1
	
	path := fmt.Sprintf("images/%d-%s",userID,file.Filename)
	err = c.SaveUploadedFile(file,path)
	if err != nil{
		data := gin.H{"is_uploaded":false}
		response := helper.APIResponse("Failed to upload avatar",http.StatusBadRequest,"error",data)
		
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_,err = h.userService.SaveAvatar(userID,path)
	if err != nil{
		data := gin.H{"is_uploaded":false}
		response := helper.APIResponse("Failed to upload avatar",http.StatusBadRequest,"error",data)
		
		c.JSON(http.StatusBadRequest, response)
		return

	}
	
	data := gin.H{"is_uploaded":true}
	response := helper.APIResponse("Success to upload avatar",http.StatusOK,"success",data)
	c.JSON(http.StatusOK,response)
}