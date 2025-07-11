package controllers

import (
	"go_fiber/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllProducts(c *fiber.Ctx)error{

	db := c.Locals("db").(*gorm.DB)
	var products []models.Product
	db.Find(&products)
	return c.JSON(products)
}

func CreateProduct(c *fiber.Ctx)error{
	db := c.Locals("db").(*gorm.DB)
	userID := c.Locals("user_id").(uint) 

	var product models.Product
	if err := c.BodyParser(&product); err != nil{
		return err
	}
	product.UserID = userID
	if err := db.Create(&product).Error; err != nil{
		return err
	}
	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	userID := c.Locals("user_id").(uint)
	role := c.Locals("role").(string)

	id, _ := strconv.Atoi(c.Params("id"))

	var product models.Product
	if err := db.First(&product, id).Error; err != nil {
		return fiber.ErrNotFound
	}

	if product.UserID != userID && role != "admin" {
		return fiber.ErrUnauthorized
	}

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	db.Save(&product)
	return c.JSON(product)
}


func DeleteProduct(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	userID := c.Locals("user_id").(uint)
	role := c.Locals("role").(string)

	id, _ := strconv.Atoi(c.Params("id"))

	var product models.Product
	if err := db.First(&product, id).Error; err != nil {
		return fiber.ErrNotFound
	}

	if product.UserID != userID && role != "admin" {
		return fiber.ErrUnauthorized
	}

	db.Delete(&product)
	return c.JSON(fiber.Map{"message": "Product deleted"})
}