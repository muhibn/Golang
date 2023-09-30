package main

import (
	"fmt"
	"server/database"
	"server/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()
	db, err := database.Connection()
	if err != nil {
		fmt.Println(err)
		return
	} else if db != nil {
		fmt.Println(db, "Data base connect is successful ")
	}

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Routes(app)

	app.Listen(":8000")
}
