package app

import (
	"github.com/faruqii/Midterm-Exam-EAI/internal/config"
	"github.com/faruqii/Midterm-Exam-EAI/internal/repositories"
	"github.com/faruqii/Midterm-Exam-EAI/internal/routes"
	"github.com/faruqii/Midterm-Exam-EAI/internal/services"
	"github.com/gin-gonic/gin"
)

func StartApplication() {

	// initialize gin
	app := gin.Default()

	// initialize db
	db, err := config.Connect()
	if err != nil {
		panic(err)
	}

	// initialize repositories
	userRepository := repositories.NewUserRepository(db)

	// initialize services
	userService := services.NewUserService(userRepository)

	// initialize routes
	apiEndpoint := app.Group("/api")
	userRouter := apiEndpoint.Group("/user")
	routes.SetUpUserRoutes(userRouter, userService)

	// start the server

	app.Run(":3000")
}
