package controllers

import (
	"go_fiber/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	var input models.User
	if err := c.BodyParser(&input); err != nil {
		return err
	}

	hashed,_:= bcrypt.GenerateFromPassword([]byte(input.Password),14)
	user := models.User{
		Username: input.Username,
		Password: string(hashed),
		Role:    "user",
	}
	if err:= db.Create(&user).Error;err != nil{
		return c.Status(400).JSON(fiber.Map{
			"error": "Failed to register user",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success register user",
	}) 
}

func Login(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	var input models.User
	if err := c.BodyParser(&input); err != nil {
		return err
	}

	var user models.User
	if err := db.Where("username = ?", input.Username).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "User tidak ditemukan"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Password salah"})
	}

	// JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"token": t})
}