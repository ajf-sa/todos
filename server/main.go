package main

import (
	"log"

	"github.com/alfuhigi/todos/server/db"
	"github.com/alfuhigi/todos/server/handlers"
	"github.com/alfuhigi/todos/server/providers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	repo := db.NewRepo(providers.Connect())
	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())
	hd := handlers.NewHandlers(repo)
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Todos App"))
	})

	app.Get("/todos", hd.ListTodos)
	app.Post("/todo/create", hd.SetTodo)

	app.Listen(":3000")

}
