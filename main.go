package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pfirulo2022/fiber-crud-api/database"
)

func main() {
	database.ConnectDB()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))

	log.Fatal(app.Listen(":8000"))
}
