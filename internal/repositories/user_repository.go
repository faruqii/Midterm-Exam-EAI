package repositories

import (
	"github.com/faruqii/Midterm-Exam-EAI/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
	Update(user *domain.User) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindUserByToken(token string) (*domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Insert(user *domain.User) (*domain.User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Update(user *domain.User) (*domain.User, error) {
	err := r.db.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindUserByToken(token string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("token = ?", token).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
