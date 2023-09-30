package routes

import (
	"server/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Post("/api/login", controllers.Logout)
	app.Post("/api/logout", controllers.User)
}
