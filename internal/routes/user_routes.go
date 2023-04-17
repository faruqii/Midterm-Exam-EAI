package routes

import (
	"github.com/faruqii/Midterm-Exam-EAI/internal/controllers"
	"github.com/faruqii/Midterm-Exam-EAI/internal/middleware"
	"github.com/faruqii/Midterm-Exam-EAI/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetUpUserRoutes(router fiber.Router, userService services.UserService) {
	userController := controllers.NewUserController(userService)
	router.Post("/register", userController.Register)
	router.Post("/login", userController.Login)

	user := router.Group("/user").Use(middleware.UserAuthentication(middleware.AuthConfig{
		Unauthorized: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	}))

	user.Get("/profile", userController.GetUserProfile)

	balance := user.Group("/balance")
	balance.Post("/topup", userController.TopUp)
	user.Get("", userController.GetBalance)

}
