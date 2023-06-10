package main

import (
	"fmt"
	"github.com/qazxcdswe123/golang-todo-mvc/app/dal"
	"github.com/qazxcdswe123/golang-todo-mvc/app/routes"
	"github.com/qazxcdswe123/golang-todo-mvc/config"
	"github.com/qazxcdswe123/golang-todo-mvc/config/database"
	"github.com/qazxcdswe123/golang-todo-mvc/utils"

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
