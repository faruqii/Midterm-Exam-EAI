package services

import (
	"net/http"

	"github.com/faruqii/Midterm-Exam-EAI/internal/config"
	"github.com/faruqii/Midterm-Exam-EAI/internal/domain"
	"github.com/faruqii/Midterm-Exam-EAI/internal/repositories"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(user *domain.User) (*domain.User, error)
	BeforeCreate(user *domain.User) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *userService {
	return &userService{userRepository: userRepository}
}

func (s *userService) Register(user *domain.User) (*domain.User, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewUserRepository(conn)

	// check if email already exist
	_, err = repo.FindByEmail(user.Email)

	if err == nil {
		return nil, &ErrorMessage{
			Message: "Email already exist",
			Code:    http.StatusBadRequest,
		}
	}

	// Validate the role of the user
	_, err = repo.FindRoleByName(user.Role)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Invalid role",
			Code:    http.StatusBadRequest,
		}
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to hash password",
			Code:    http.StatusInternalServerError,
		}
	}

	user.Password = string(hashedPassword)

	// Insert the user to database
	user, err = repo.Insert(user)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to register",
			Code:    http.StatusInternalServerError,
		}
	}

	return user, nil
}

func (s *userService) BeforeCreate(user *domain.User) error {
	user.ID = uuid.NewString()
	return nil
}
