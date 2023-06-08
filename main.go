package main

import (
	"fmt"
	"numtostr/gotodo/app/dal"
	"numtostr/gotodo/app/routes"
	"numtostr/gotodo/config"
	"numtostr/gotodo/config/database"
	"numtostr/gotodo/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	database.Migrate(&dal.User{}, &dal.Todo{})

	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	app.Use(logger.New())
	app.Use(cors.New())

	routes.AuthRoutes(app)
	routes.TodoRoutes(app)

	app.Listen(fmt.Sprintf("localhost:%v", config.PORT))
}
