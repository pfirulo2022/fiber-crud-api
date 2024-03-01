package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pfirulo2022/fiber-crud-api/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/hello", handlers.Hello)
	app.Get("/get", handlers.GetTasks)
	app.Post("/post", handlers.CreateTask)
	app.Get("/get/:id", handlers.GetTask)
	app.Put("/update/:id", handlers.UpdateTask)
	app.Delete("/delete/:id", handlers.DeleteTask)

}
