package main

import (
	"github.com/gofiber/fiber/v2"
	"book-api/database"
	"book-api/handlers"
)

func main() {
	database.Connect()
	app := fiber.New()

	app.Get("/books", handlers.GetBooks)
	app.Get("/books/:id", handlers.GetBook)
	app.Post("/books", handlers.CreateBook)
	app.Put("/books/:id", handlers.UpdateBook)
	app.Delete("/books/:id", handlers.DeleteBook)

	app.Listen(":3000")
}
