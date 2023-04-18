package controllers

import (
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
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success add product",
		"data":    productResponse,
	})
}

func (p *ProductController) UpdateProduct(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	req := dto.ProductRequest{}

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

	product, err = p.productService.UpdateProduct(product)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success update product",
		"data":    product,
	})
}

func (p *ProductController) GetAllProducts(ctx *fiber.Ctx) (err error) {
	products, err := p.productService.GetAllProduct()

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success get all products",
		"data":    products,
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

// Find Product
// using query string to find product
// example: /products?name=product1
// also can use multiple query string
// example: /products?name=product1&category_name=category1
func (p *ProductController) FindProduct(ctx *fiber.Ctx) (err error) {
	name := ctx.Query("name")
	categoryName := ctx.Query("category_name")

	var products []domain.Product

	if name != "" {
		products, err = p.productService.FindProductByName(name)
	} else if categoryName != "" {
		products, err = p.productService.FinProductByCategoryName(categoryName)
	} else {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "query string not found",
		})
	}

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// loop all products
	// and get category name
	response := []dto.ProductResponse{}

	for _, product := range products {
		category, err := p.productService.GetCategoryByID(product.CategoryID)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		response = append(response, dto.ProductResponse{
			Name:         product.Name,
			Description:  product.Description,
			Price:        product.Price,
			Stock:        product.Stock,
			CategoryName: category.Name,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success find product",
		"status":  "success",
		"data":    response,
	})

}
