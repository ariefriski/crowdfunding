package handler

import (
	"crowdfunding/helper"
	"crowdfunding/user"
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

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Account failed registered", http.StatusUnprocessableEntity, "error", err.Error())
		c.JSON(http.StatusBadGateway, response)
		return
	}

	formatter := user.FormatUser(newUser, "tokencontoh")

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context) {
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
