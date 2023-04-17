package middleware

import (
	"github.com/faruqii/Midterm-Exam-EAI/internal/config"
	"github.com/faruqii/Midterm-Exam-EAI/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type AuthConfig struct {
	Filter       func(*fiber.Ctx) error
	Unauthorized fiber.Handler
}

func UserAuthentication(c AuthConfig) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		header := ctx.GetReqHeaders()

		if _, ok := header["Authorization"]; !ok {
			return c.Unauthorized(ctx)
		}

		// check user token is valid
		userToken := domain.Token{}

		err := config.DB.Where("token = ?", header["Authorization"]).First(&userToken).Error

		if err != nil {
			return c.Unauthorized(ctx)
		}

		if userToken.Type != "user" {
			return c.Unauthorized(ctx)
		}

		ctx.Locals("user", userToken)

		return ctx.Next()
	}
}

func AdminAuthentication(c AuthConfig) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		header := ctx.GetReqHeaders()

		if _, ok := header["Authorization"]; !ok {
			return c.Unauthorized(ctx)
		}

		// check admin token is valid
		var adminToken domain.Token

		err := config.DB.Where("token = ?", header["Authorization"]).First(&adminToken).Error

		if err != nil {
			return c.Unauthorized(ctx)
		}

		if adminToken.Type != "admin" {
			return c.Unauthorized(ctx)
		}

		ctx.Locals("admin", adminToken.Token)

		return ctx.Next()

	}
}
