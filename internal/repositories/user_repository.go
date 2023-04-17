package repositories

import (
	"github.com/faruqii/Midterm-Exam-EAI/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
	Update(user *domain.User) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindByID(id string) (*domain.User, error)
	FindUserByToken(token string) (string, error)
	FindRoleByName(name string) (*domain.Role, error)
	CreateUserBalance(userBalance *domain.UserBalance) (*domain.UserBalance, error)
	UpdateUserBalance(userBalance *domain.UserBalance) (*domain.UserBalance, error)
	GetUserBalance(UserID string) (*domain.UserBalance, error)
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

func (r *userRepository) FindByID(id string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindUserByToken(token string) (string, error) {
	var user domain.Token

	err := r.db.Where("token = ?", token).First(&user).Error

	if err != nil {
		return "", err
	}

	return user.UserID, nil
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

func (r *userRepository) CreateUserBalance(userBalance *domain.UserBalance) (*domain.UserBalance, error) {
	err := r.db.Create(userBalance).Error
	if err != nil {
		return nil, err
	}
	return userBalance, nil
}

func (r *userRepository) UpdateUserBalance(userBalance *domain.UserBalance) (*domain.UserBalance, error) {
	err := r.db.Save(userBalance).Error
	if err != nil {
		return nil, err
	}
	return userBalance, nil
}

func (r *userRepository) GetUserBalance(UserID string) (*domain.UserBalance, error) {
	var userBalance domain.UserBalance
	err := r.db.Where("user_id = ?", UserID).First(&userBalance).Error
	if err != nil {
		return nil, err
	}

	return &userBalance, nil
}
