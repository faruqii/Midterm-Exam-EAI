package repositories

import (
	"github.com/faruqii/Midterm-Exam-EAI/internal/domain"
	"gorm.io/gorm"
)

type ProductsRepository interface {
	Insert(product *domain.Product) (*domain.Product, error)
	Update(product *domain.Product) (*domain.Product, error)
	Delete(id string) error
	FindAll() ([]domain.Product, error)
	FindByID(id string) (*domain.Product, error)
	FindByCategoryName(name string) ([]domain.Product, error)
	FindProductByName(name string) ([]domain.Product, error)
	GetAllCategories() ([]domain.Categories, error)
	GetCategoriesByID(id string) (*domain.Categories, error)
}

type productsRepository struct {
	db *gorm.DB
}

func NewProductsRepository(db *gorm.DB) *productsRepository {
	return &productsRepository{db: db}
}

func (r *productsRepository) Insert(product *domain.Product) (*domain.Product, error) {
	err := r.db.Create(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productsRepository) Update(product *domain.Product) (*domain.Product, error) {
	err := r.db.Save(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productsRepository) Delete(id string) error {
	var product domain.Product
	err := r.db.Where("id = ?", id).Delete(&product).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *productsRepository) FindAll() ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.Preload("Category").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productsRepository) FindByID(id string) (*domain.Product, error) {
	var product domain.Product
	err := r.db.Preload("Category").Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productsRepository) FindByCategoryName(name string) ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.Joins("Category").Where("categories.name = ?", name).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productsRepository) FindProductByName(name string) ([]domain.Product, error) {
	var products []domain.Product

	err := r.db.Preload("Category").Where("name LIKE ?", "%"+name+"%").Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productsRepository) GetAllCategories() ([]domain.Categories, error) {
	var categories []domain.Categories
	err := r.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *productsRepository) GetCategoriesByID(id string) (*domain.Categories, error) {
	var category domain.Categories
	err := r.db.Where("id = ?", id).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}
