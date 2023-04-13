package repositories

import (
	"github.com/faruqii/Midterm-Exam-EAI/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
	Update(user *domain.User) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindUserByToken(token string) (*domain.User, error)
	FindRoleByName(name string) (*domain.Role, error)
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

func (r *userRepository) BeforeCreate(scope *gorm.DB) error {
	scope.Statement.SetColumn("ID", uuid.NewString())
	return nil
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

func (r *userRepository) FindRoleByName(name string) (*domain.Role, error) {
	// default role is user
	if name == "" {
		name = "user"
	}

	var role domain.Role

	err := r.db.Where("name = ?", name).First(&role).Error

	if err != nil {
		return nil, err
	}

	return &role, nil
}
