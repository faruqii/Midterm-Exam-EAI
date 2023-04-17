package services

import (
	"net/http"

	"github.com/faruqii/Midterm-Exam-EAI/internal/config"
	"github.com/faruqii/Midterm-Exam-EAI/internal/domain"
	"github.com/faruqii/Midterm-Exam-EAI/internal/repositories"
)

type ProductService interface {
	AddProduct(product *domain.Product) (*domain.Product, error)
	UpdateProduct(product *domain.Product) (*domain.Product, error)
	GetAllProduct() ([]domain.Product, error)
	GetProductByID(id string) (*domain.Product, error)
	GetProductByCategoryName(name string) ([]domain.Product, error)
}

type productService struct {
	productRepository repositories.ProductsRepository
}

func NewProductService(productRepository repositories.ProductsRepository) *productService {
	return &productService{productRepository: productRepository}
}

func (s *productService) AddProduct(product *domain.Product) (*domain.Product, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewProductsRepository(conn)

	// check if product already exist
	_, err = repo.FindProductByName(product.Name)

	if err == nil {
		return nil, &ErrorMessage{
			Message: "Product already exist",
			Code:    http.StatusBadRequest,
		}
	}

	// Validate the category of the product
	_, err = repo.FindByCategoryName(product.Category.Name)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Invalid category",
			Code:    http.StatusBadRequest,
		}
	}

	return repo.Insert(product)

}

func (s *productService) UpdateProduct(product *domain.Product) (*domain.Product, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewProductsRepository(conn)

	// check if product already exist
	_, err = repo.FindProductByName(product.Name)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Product not found",
			Code:    http.StatusBadRequest,
		}
	}

	// Validate the category of the product
	_, err = repo.FindByCategoryName(product.Category.Name)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Invalid category",
			Code:    http.StatusBadRequest,
		}
	}

	return repo.Update(product)

}

func (s *productService) GetAllProduct() ([]domain.Product, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewProductsRepository(conn)

	return repo.FindAll()

}

func (s *productService) GetProductByID(id string) (*domain.Product, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewProductsRepository(conn)

	return repo.FindByID(id)

}

func (s *productService) GetProductByCategoryName(name string) ([]domain.Product, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewProductsRepository(conn)

	return repo.FindByCategoryName(name)

}

