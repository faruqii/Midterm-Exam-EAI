package controllers

import (
	"time"

	"github.com/faruqii/Midterm-Exam-EAI/internal/domain"
	"github.com/faruqii/Midterm-Exam-EAI/internal/dto"
	"github.com/gofiber/fiber/v2"
)

func (p *ProductController) GetCategories(ctx *fiber.Ctx) (err error) {
	categories, err := p.productService.GetAllCategories()

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success get categories",
		"data":    categories,
	})
}

func (p *ProductController) AddProduct(ctx *fiber.Ctx) (err error) {
	req := dto.ProductRequest{}

	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	product := &domain.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CategoryID:  req.CategoryID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	product, err = p.productService.AddProduct(product)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	product, err = p.productService.GetProductByID(product.ID)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	productResponse := dto.ProductResponse{
		ID:           product.ID,
		Name:         product.Name,
		Description:  product.Description,
		Price:        product.Price,
		Stock:        product.Stock,
		CategoryName: product.Category.Name,
		CreatedAt:    product.CreatedAt,
		UpdatedAt:    product.UpdatedAt,
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success add product",
		"data":    productResponse,
	})
}

func (p *ProductController) UpdateProduct(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	req := dto.ProductUpdateRequest{}

	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	product, err := p.productService.GetProductByID(id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.Stock = req.Stock
	product.CategoryID = req.CategoryID
	product.UpdatedAt = time.Now()

	product, err = p.productService.UpdateProduct(product)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.ProductResponse{
		ID:           product.ID,
		Name:         product.Name,
		Description:  product.Description,
		Price:        product.Price,
		Stock:        product.Stock,
		CategoryName: product.Category.Name,
		CreatedAt:    product.CreatedAt,
		UpdatedAt:    product.UpdatedAt,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success update product",
		"data":    response,
	})
}

func (p *ProductController) GetAllProducts(ctx *fiber.Ctx) (err error) {
	products, err := p.productService.GetAllProduct()

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// loop all products then add to response
	response := []dto.ProductResponse{}

	for _, product := range products {
		response = append(response, dto.ProductResponse{
			ID:           product.ID,
			Name:         product.Name,
			Description:  product.Description,
			Price:        product.Price,
			Stock:        product.Stock,
			CategoryName: product.Category.Name,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success get all products",
		"data":    response,
	})
}

func (p *ProductController) DeleteProduct(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	err = p.productService.DeleteProduct(id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success delete product",
	})
}

// FindProduct find product by name or category name
// using query string like this: product?name=tangkelek or product?category_name=food or both of them like this: product?name=tangkelek&category_name=food
func (p *ProductController) FindProduct(ctx *fiber.Ctx) (err error) {
	name := ctx.Query("name")
	categoryName := ctx.Query("category_name")

	// if both of query string is empty
	if name == "" && categoryName == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "query string name or category_name is required",
		})
	}

	response := []dto.ProductResponse{}

	if name != "" {
		products, err := p.productService.FindProductByName(name)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// loop all products then add to response

		for _, product := range products {
			response = append(response, dto.ProductResponse{
				ID:           product.ID,
				Name:         product.Name,
				Description:  product.Description,
				Price:        product.Price,
				Stock:        product.Stock,
				CategoryName: product.Category.Name,
			})
		}
	}

	if categoryName != "" {
		products, err := p.productService.FinProductByCategoryName(categoryName)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// loop all products then add to response

		for _, product := range products {
			response = append(response, dto.ProductResponse{
				ID:           product.ID,
				Name:         product.Name,
				Description:  product.Description,
				Price:        product.Price,
				Stock:        product.Stock,
				CategoryName: product.Category.Name,
			})
		}

	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success find product",
		"data":    response,
	})
}
