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

func CreatePost(c *fiber.Ctx) error {
	posts := new(models.Post)
	if err := c.BodyParser(&posts); err != nil {
		return err
	}

	//validasi request
	validate := validator.New()
	errValidate := validate.Struct(posts)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error": errValidate.Error(),
		})
	}

	if len(posts.TagsID) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"err": "tags_id is required",
		})
	}


	newPost := models.Post{
		Title : posts.Title,
		Description: posts.Description,
		UserID: posts.UserID,
		TagsID: posts.TagsID,
	}

	errCreatePost := database.Db.Debug().Create(&newPost).Error
	if errCreatePost != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "gagal menyimpan data" + errCreatePost.Error(),
		})
	}

	for _, tagID := range newPost.TagsID {
		postTag := new(models.PostTag)
		postTag.PostID = newPost.ID
		postTag.TagID = tagID
		database.Db.Create(&postTag)
	}

	//if succeed
	return c.JSON(fiber.Map{
		"message" : "success",
		"posts": newPost,
	})
}

func GetAllPost(c *fiber.Ctx) error  {
	var posts []models.PostResponseWithTags
	result := database.Db.Preload("User").Preload("Tag").Find(&posts)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return c.JSON(fiber.Map{
		"posts": posts,
	})
}
