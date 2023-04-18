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
	DeleteProduct(id string) error
	GetAllProduct() ([]domain.Product, error)
	GetProductByID(id string) (*domain.Product, error)
	GetProductByCategoryName(name string) ([]domain.Product, error)
	GetAllCategories() ([]domain.Categories, error)
	FindProductByName(name string) ([]domain.Product, error)
	FinProductByCategoryName(category string) ([]domain.Product, error)
	GetCategoryByID(id string) (*domain.Categories, error)
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

	// check category exist or not by category id
	_, err = repo.GetCategoriesByID(product.Category.ID)


	if err == nil {
		return nil, &ErrorMessage{
			Message: "Product already exist",
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

func (s *productService) DeleteProduct(id string) error {
	conn, err := config.Connect()

	if err != nil {
		return &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewProductsRepository(conn)

	// check if product already exist

	_, err = repo.FindByID(id)

	if err != nil {
		return &ErrorMessage{
			Message: "Product not found",
			Code:    http.StatusBadRequest,
		}
	}

	return repo.Delete(id)
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

func (s *productService) GetAllCategories() ([]domain.Categories, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewProductsRepository(conn)

	return repo.GetAllCategories()
}

func (s *productService) FindProductByName(name string) ([]domain.Product, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewProductsRepository(conn)

	return repo.FindProductByName(name)
}

func (s *productService) FinProductByCategoryName(category string) ([]domain.Product, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewProductsRepository(conn)

	return repo.FindByCategoryName(category)
}

func (s *productService) GetCategoryByID(id string) (*domain.Categories, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewProductsRepository(conn)

	return repo.GetCategoriesByID(id)
}
