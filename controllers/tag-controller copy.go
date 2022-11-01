package controllers

import (
	//"strconv"
	"learning/golang-gormdatabase/database"
	"learning/golang-gormdatabase/models"

	//"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreateTag(c *fiber.Ctx) error {
	tags := new(models.Tag)
	if err := c.BodyParser(&tags); err != nil {
		return err
	}

	//validasi request
	validate := validator.New()
	errValidate := validate.Struct(tags)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error": errValidate.Error(),
		})
	}

	newTag := models.Tag{
		Name : tags.Name,
	}

	errCreateTag := database.Db.Create(&newTag).Error
	if errCreateTag != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "gagal menyimpan data" + errCreateTag.Error(),
		})
	}
	//if succeed
	return c.JSON(fiber.Map{
		"message" : "success",
		"tags": newTag,
	})
}

func GetAllTag(c *fiber.Ctx) error  {
	var tags []models.TagResponseWithPost
	result := database.Db.Debug().Preload("Posts").Find(&tags)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return c.JSON(fiber.Map{
		"tags": tags,
	})
}
