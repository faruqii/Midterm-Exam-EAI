package controllers

import "github.com/faruqii/Midterm-Exam-EAI/internal/services"

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}
