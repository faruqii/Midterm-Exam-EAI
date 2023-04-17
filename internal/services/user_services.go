package services

import (
	"net/http"

	"github.com/faruqii/Midterm-Exam-EAI/internal/config"
	"github.com/faruqii/Midterm-Exam-EAI/internal/domain"
	"github.com/faruqii/Midterm-Exam-EAI/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(user *domain.User) (*domain.User, error)
	Login(email, password string) (*domain.User, error)
	GetUserProfile(userID string) (*domain.User, error)
	FindRoleByName(name string) (*domain.Role, error)
	FindUserByToken(token string) (*domain.User, error)
	GetUserBalance(UserID string) (*domain.UserBalance, error)
	AddBalance(userBalance *domain.UserBalance) (*domain.UserBalance, error)
	UpdateBalance(userBalance *domain.UserBalance) (*domain.UserBalance, error)
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

func (s *userService) Login(email, password string) (*domain.User, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewUserRepository(conn)

	user, err := repo.FindByEmail(email)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "User not found",
			Code:    http.StatusNotFound,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Wrong password",
			Code:    http.StatusUnauthorized,
		}
	}

	return user, nil
}

func (s *userService) GetUserProfile(userID string) (*domain.User, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewUserRepository(conn)

	user, err := repo.FindByID(userID)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "User not found",
			Code:    http.StatusNotFound,
		}
	}

	return user, nil
}

func (s *userService) FindRoleByName(name string) (*domain.Role, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewUserRepository(conn)

	role, err := repo.FindRoleByName(name)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Role not found",
			Code:    http.StatusNotFound,
		}
	}

	return role, nil
}

func (s *userService) FindUserByToken(token string) (*domain.User, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewUserRepository(conn)

	userID, err := repo.FindUserByToken(token)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Token not found",
			Code:    http.StatusNotFound,
		}
	}

	user, err := repo.FindByID(userID)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "User not found",
			Code:    http.StatusNotFound,
		}
	}

	return user, nil
}

func (s *userService) GetUserBalance(UserID string) (*domain.UserBalance, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewUserRepository(conn)

	userBalance, err := repo.GetUserBalance(UserID)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "User not found",
			Code:    http.StatusNotFound,
		}
	}

	return userBalance, nil
}

func (s *userService) AddBalance(userBalance *domain.UserBalance) (*domain.UserBalance, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewUserRepository(conn)

	existingUserBalance, err := repo.GetUserBalance(userBalance.UserID)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "User not found",
			Code:    http.StatusNotFound,
		}
	}

	if existingUserBalance == nil {
		// if user balance not exist, create new user balance
		newUserBalance, err := repo.CreateUserBalance(userBalance)

		if err != nil {
			return nil, &ErrorMessage{
				Message: "Failed to create user balance",
				Code:    http.StatusInternalServerError,
			}
		}

		return newUserBalance, nil
	} else {
		return nil, &ErrorMessage{
			Message: "User balance already exist",
			Code:    http.StatusBadRequest,
		}
	}
}

func (s *userService) UpdateBalance(userBalance *domain.UserBalance) (*domain.UserBalance, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewUserRepository(conn)

	updatedUserBalance, err := repo.UpdateUserBalance(userBalance)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to update user balance",
			Code:    http.StatusInternalServerError,
		}
	}

	return updatedUserBalance, nil
}
