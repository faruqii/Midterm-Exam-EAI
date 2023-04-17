package controllers

import (
	"net/http"
	"time"

	"github.com/faruqii/Midterm-Exam-EAI/internal/config"
	"github.com/faruqii/Midterm-Exam-EAI/internal/domain"
	"github.com/faruqii/Midterm-Exam-EAI/internal/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (c *UserController) Register(ctx *fiber.Ctx) error {
	var registerRequest dto.RegisterRequest

	if err := ctx.BodyParser(&registerRequest); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	role, err := c.userService.FindRoleByName("user")

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user := &domain.User{
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
		Phone:    registerRequest.Phone,
		Role:     role.Name,
	}

	// save user to database
	user, err = c.userService.Register(user)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// return response
	response := dto.RegisterResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
		Role:  user.Role,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "User registered successfully",
		"data":    response,
	})

}

func (c *UserController) Login(ctx *fiber.Ctx) (err error) {
	req := dto.LoginRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := c.userService.Login(req.Email, req.Password)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, err := CreateUserToken(user)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// set cookie
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	// return response
	response := dto.LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
		Role:  user.Role,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "User logged in successfully",
		"data":    response,
		"token":   token,
	})

}

func CreateUserToken(user *domain.User) (string, error) {
	claims := dto.Claims{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    user.Name,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString([]byte("secret"))

	userToken := domain.Token{
		UserID: user.ID,
		Token:  signedToken,
		Type:   user.Role,
	}

	var existingUserToken domain.Token

	err := config.DB.Where("user_id = ?", user.ID).First(&existingUserToken).Error

	if err != nil {
		err = config.DB.Create(&userToken).Error
	} else {
		err = config.DB.Model(&existingUserToken).Updates(&userToken).Error
	}

	return signedToken, err
}

func (c *UserController) GetUserProfile(ctx *fiber.Ctx) (err error) {
	user := ctx.Locals("user").(domain.Token)

	userProfile, err := c.userService.GetUserProfile(user.UserID)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// return response
	response := dto.ProfileResponse{
		ID:    userProfile.ID,
		Name:  userProfile.Name,
		Email: userProfile.Email,
		Phone: userProfile.Phone,
		Role:  userProfile.Role,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "User profile fetched successfully",
		"data":    response,
	})
}
