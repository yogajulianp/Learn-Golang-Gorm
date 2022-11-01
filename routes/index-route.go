package routes

import (
	"github.com/gofiber/fiber/v2"
	"learning/golang-gormdatabase/controllers"
)

func RouteInit(app *fiber.App) {
	app.Get("/users", controllers.UserGetAll)
	app.Post("/users", controllers.UserRegister)

	app.Get("/lockers", controllers.GetAllLocker)
	app.Post("/lockers", controllers.CreateLocker)

	app.Get("/post", controllers.GetAllPost)
	app.Post("/post", controllers.CreatePost)

	
}