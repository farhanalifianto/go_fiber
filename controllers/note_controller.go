package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"go_fiber/models"
)

func GetNotes(c *fiber.Ctx)error{
	db := c.Locals("db").(*gorm.DB) 
	userID := c.Locals("user_id").(uint)
	role := c.Locals("role").(string)

	var notes []models.Note

	if role == "admin"{
		db.Find(&notes)
	}else{
		db.Where("user_id = ?", userID).Find(&notes)
	}
	if len(notes) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Masih Kosong",
		})
	}
	return c.JSON(notes)
}

func CreateNote(c *fiber.Ctx)error{
	db := c.Locals("db").(*gorm.DB)
	userID := c.Locals("user_id").(uint)

	var note models.Note
	if err := c.BodyParser(&note); err != nil{
		return err
	}
	note.UserID = userID
	if err := db.Create(&note).Error; err != nil{
		return err
	}
	return c.JSON(note)
}

func UpdateNote(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	id := c.Params("id")
	userID := c.Locals("user_id").(uint)
	role := c.Locals("role").(string)

	var note models.Note
	if err := db.First(&note, id).Error; err != nil {
		return fiber.ErrNotFound
	}

	if note.UserID != userID && role != "admin" {
		return fiber.ErrUnauthorized
	}

	if err := c.BodyParser(&note); err != nil {
		return err
	}
	db.Save(&note)

	return c.JSON(note)
}

func DeleteNote(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	id := c.Params("id")
	userID := c.Locals("user_id").(uint)
	role := c.Locals("role").(string)

	var note models.Note
	if err := db.First(&note, id).Error; err != nil {
		return fiber.ErrNotFound
	}

	if note.UserID != userID && role != "admin" {
		return fiber.ErrUnauthorized
	}

	db.Delete(&note)
	return c.JSON(fiber.Map{"message": "Note deleted"})
}

func ToggleFavorite(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	noteID, _ := strconv.Atoi(c.Params("id"))
	userID := c.Locals("user_id").(uint)

	var user models.User
	if err := db.Preload("Favorites").First(&user, userID).Error; err != nil {
		return err
	}

	var note models.Note
	if err := db.First(&note, noteID).Error; err != nil {
		return fiber.ErrNotFound
	}

	// Check if already favorite
	found := false
	for _, fav := range user.Favorites {
		if fav.ID == uint(noteID) {
			found = true
			break
		}
	}

	if found {
		db.Model(&user).Association("Favorites").Delete(&note)
		return c.JSON(fiber.Map{"message": "Removed from favorites"})
	} else {
		db.Model(&user).Association("Favorites").Append(&note)
		return c.JSON(fiber.Map{"message": "Added to favorites"})
	}
}

func GetFavoriteNotes(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	userID := c.Locals("user_id").(uint)

	var user models.User
	if err := db.Preload("Favorites").First(&user, userID).Error; err != nil {
		return err
	}

	return c.JSON(user.Favorites)
}