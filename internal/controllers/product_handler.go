package controllers

import "github.com/faruqii/Midterm-Exam-EAI/internal/services"

type ProductController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{productService: productService}
}
