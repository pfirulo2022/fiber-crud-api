package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pfirulo2022/fiber-crud-api/database"
	"github.com/pfirulo2022/fiber-crud-api/models"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello from agust!")
}

// Devuelve una sola tarea
func GetTask(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var task models.Task
	db.Find(&task, id)
	return c.JSON(task)

}

// Devuelve todas las tareas
func GetTasks(c *fiber.Ctx) error {
	db := database.DB
	var tasks []models.Task
	db.Find(&tasks)
	return c.JSON(tasks)

}

// Crea una tarea
func CreateTask(c *fiber.Ctx) error {
	db := database.DB
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(400).JSON(err)
	}
	db.Create(&task)
	return c.JSON(task)

}

// Actualizar una tarea
func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var task models.Task

	if err := db.First(&task, id).Error; err != nil {
		return c.Status(404).JSON(err)
	}

	updateTask := new(models.Task)

	if err := c.BodyParser(updateTask); err != nil {
		return c.Status(400).JSON(err)
	}

	task.Body = updateTask.Body
	db.Save(&task)

	return c.JSON(task)

}

// Borrar una tarea
func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var task models.Task
	if err := db.First(&task, id).Error; err != nil {
		return c.Status(404).JSON(err)
	}
	db.Delete(&task)
	return c.JSON(fiber.Map{"status": "success", "message": "tarea eliminada correctamente"})
}
