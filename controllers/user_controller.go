package controllers

import (
	"go_fiber/dto"
	"go_fiber/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllUsers(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	var users []models.User
	if err := db.Select("id", "username", "role").Find(&users).Error; err != nil {
		return err
	}

	var responses []dto.UserResponse
	for _, user := range users {
		responses = append(responses, dto.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Role:     user.Role,
		})
	}
	

	return c.JSON(responses)
}

