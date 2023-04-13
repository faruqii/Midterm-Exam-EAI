package controllers

import (
	"net/http"

	"github.com/faruqii/Midterm-Exam-EAI/internal/domain"
	"github.com/faruqii/Midterm-Exam-EAI/internal/dto"
	"github.com/faruqii/Midterm-Exam-EAI/internal/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) Register(ctx *gin.Context) {
	var registerRequest dto.RegisterRequest
	err := ctx.ShouldBindJSON(&registerRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &domain.User{
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Phone:    registerRequest.Phone,
		Password: registerRequest.Password,
		Role:     "user",
	}

	createdUser, err := c.userService.Register(user)
	if err != nil {
		if errMsg, ok := err.(*services.ErrorMessage); ok {
			ctx.JSON(errMsg.Code, gin.H{"error": errMsg.Message})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register"})
		return
	}

	registerResponse := &dto.RegisterResponse{
		ID:       createdUser.ID,
		Name:     createdUser.Name,
		Email:    createdUser.Email,
		Phone:    createdUser.Phone,
		Role:     createdUser.Role,
	}

	response := gin.H{
		"status":  "success",
		"message": "User registered successfully",
		"data":    registerResponse,
	}

	ctx.JSON(http.StatusOK, response)
}
