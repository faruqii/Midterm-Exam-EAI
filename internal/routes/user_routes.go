package routes

import (
	"github.com/faruqii/Midterm-Exam-EAI/internal/controllers"
	"github.com/faruqii/Midterm-Exam-EAI/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetUpUserRoutes(router fiber.Router, userService services.UserService) {
	userController := controllers.NewUserController(userService)
	router.Post("/register", userController.Register)
	router.Post("/login", userController.Login)
}
