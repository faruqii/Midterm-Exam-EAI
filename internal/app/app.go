package app

import (
	"os"

	"github.com/faruqii/Midterm-Exam-EAI/internal/config"
	"github.com/faruqii/Midterm-Exam-EAI/internal/repositories"
	"github.com/faruqii/Midterm-Exam-EAI/internal/routes"
	"github.com/faruqii/Midterm-Exam-EAI/internal/services"
	"github.com/gofiber/fiber/v2"
)

func StartApplication() {

	// initialize gin
	app := fiber.New()

	// initialize db
	db, err := config.Connect()
	if err != nil {
		panic(err)
	}

	// initialize user repositories
	userRepository := repositories.NewUserRepository(db)

	// initialize user services
	userService := services.NewUserService(userRepository)

	// initialize user routes
	apiEndpoint := app.Group("/api")
	routes.SetUpUserRoutes(apiEndpoint, userService)

	// initialize product repositories
	productRepository := repositories.NewProductsRepository(db)

	// initialize product services
	productService := services.NewProductService(productRepository)

	// initialize product routes
	routes.SetUpProductRoutes(apiEndpoint, productService)

	// start the server
	err = app.Listen(":" + os.Getenv("PORT"))

	if err != nil {
		panic(err)
	}
}
