package controllers

import (
	//"strconv"
	"learning/golang-gormdatabase/database"
	"learning/golang-gormdatabase/models"
	"learning/golang-gormdatabase/utils"

	//"learning/golang-gormdatabase/models/request"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)



func UserRegister(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	//validasi request
	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error": errValidate.Error(),
		})
	}

	newUser := models.User{
		Name    : user.Name,
		Email	: user.Email,
		Username : user.Username,
	}

	HashPassword, err := utils.HashingPassword(user.Password)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}
	newUser.Password = HashPassword

	errCreateUser := database.Db.Create(&newUser).Error
	if errCreateUser != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "gagal menyimpan data",
		})
	}
	//if succeed
	return c.JSON(fiber.Map{
		"message" : "success",
		"data": newUser,
	})
}

func UserGetAll(c *fiber.Ctx) error  {
	var users []models.User
	result := database.Db.Preload("Locker").Preload("Posts").Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return c.JSON(fiber.Map{
		"users": users,
	})
}

func UserGetById(c *fiber.Ctx) error {
	userId := c.Params("id")
	//userId,_ := strconv.Atoi(Id)
	fmt.Println(userId)

	var user models.User
	err := database.Db.Where("id = ?", userId).First(&user).Error
	
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}
	
	//if succeed
	return c.JSON(fiber.Map{
		"message" : "success",
		"data": user,
	})
}

func UserUpdate(c *fiber.Ctx) error {
	userId := c.Params("id")

	var user models.User
	err := database.Db.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	userRequest := new(models.User)
	if err := c.BodyParser(&userRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	//Update User Data  
	if userRequest.Username != "" {
		user.Username = userRequest.Username
	}
	user.Name = userRequest.Name
	user.Email = userRequest.Email
	user.Password = userRequest.Password

	errUpdate := database.Db.Save(&user).Error
	if errUpdate != nil {
		return c.Status(500).JSON(fiber.Map{
			"message" : "server error",
		})
	}

	//if succeed
	return c.JSON(fiber.Map{
		"message" : "success data has been updated",
		"data": user,
	})

}

func UserDelete(c *fiber.Ctx) error {
	userId := c.Params("id")

	var user models.User
	err := database.Db.Debug().Where("id = ?", userId).First(&user).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	errDelete := database.Db.Debug().Delete(&user).Error
	if errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"message" : "server error",
		})
	}

		//if succeed
		return c.JSON(fiber.Map{
			"message" : "user has been deleted",
		})
}