package main

import (
	"github.com/faruqii/Midterm-Exam-EAI/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	// Start the application
	app.StartApplication()
}
