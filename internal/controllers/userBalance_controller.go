package controllers

import (
	"log"

	"github.com/faruqii/Midterm-Exam-EAI/internal/domain"
	"github.com/faruqii/Midterm-Exam-EAI/internal/dto"
	"github.com/gofiber/fiber/v2"
)

func (c *UserController) TopUp(ctx *fiber.Ctx) (err error) {
	// get token from context | This is middleware
	userToken := ctx.Locals("user").(domain.Token)

	req := dto.TopupRequest{}

	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// get user id
	user, err := c.userService.FindUserByToken(userToken.Token)

	log.Println(user)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// get user balance
	userBalance, err := c.userService.GetUserBalance(user.ID)

	if err != nil {
		// create new user balance
		userBalance = &domain.UserBalance{
			UserID:  user.ID,
			Balance: req.Balance,
		}

		// save user balance to database
		userBalance, err = c.userService.AddBalance(userBalance)

		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	} else {
		// update user balance
		userBalance.Balance += req.Balance

		// save user balance to database
		userBalance, err = c.userService.UpdateBalance(userBalance)

		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	// return response
	response := dto.TopupResponse{
		UserName: user.Name,
		Balance:  userBalance.Balance,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Top up balance successfully",
		"data":    response,
	})
}
