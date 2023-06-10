package routes

import (
	"github.com/qazxcdswe123/golang-todo-mvc/app/services"
	"github.com/qazxcdswe123/golang-todo-mvc/utils/middleware"

	"github.com/gofiber/fiber/v2"
)

// TodoRoutes contains all routes relative to /todo
func TodoRoutes(app fiber.Router) {
	r := app.Group("/todo").Use(middleware.Auth)

	r.Post("/create", services.CreateTodo)
	r.Post("/search", services.SearchTodos)
	r.Get("/list", services.GetTodos)
	r.Patch("/:todoID", services.UpdateTodoTitle)
	r.Patch("/:todoID/check", services.CheckTodo)
	r.Delete("/:todoID", services.DeleteTodo)
}
