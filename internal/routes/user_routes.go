package routes

import (
	"github.com/faruqii/Midterm-Exam-EAI/internal/controllers"
	"github.com/faruqii/Midterm-Exam-EAI/internal/services"
	"github.com/gin-gonic/gin"
)

func SetUpUserRoutes(router *gin.RouterGroup, userService services.UserService) {
	userController := controllers.NewUserController(userService)
	router.POST("/register", userController.Register)
}
