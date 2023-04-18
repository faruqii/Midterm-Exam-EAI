package routes

import (
	"github.com/faruqii/Midterm-Exam-EAI/internal/controllers"
	"github.com/faruqii/Midterm-Exam-EAI/internal/middleware"
	"github.com/faruqii/Midterm-Exam-EAI/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetUpProductRoutes(router fiber.Router, productServices services.ProductService) {
	productController := controllers.NewProductController(productServices)

	// Categories
	categories := router.Group("/categories")
	categories.Get("", productController.GetCategories)

	// Products
	products := router.Group("/products").Use(middleware.AdminAuthentication(middleware.AuthConfig{
		Unauthorized: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	}))

	products.Post("", productController.AddProduct)
	products.Get("", productController.GetAllProducts)

}
