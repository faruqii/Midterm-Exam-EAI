package services

import (
	"net/http"

	"github.com/faruqii/Midterm-Exam-EAI/internal/config"
	"github.com/faruqii/Midterm-Exam-EAI/internal/domain"
	"github.com/faruqii/Midterm-Exam-EAI/internal/repositories"
)

type UserService interface {
	Insert(user *domain.User) (*domain.User, error)
	Update(user *domain.User) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindUserByToken(token string) (*domain.User, error)
	Register(user *domain.User) (*domain.User, error)
	Login(user *domain.User) (*domain.User, error)
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

	// TODO: Later hehe

	return nil, nil

}
