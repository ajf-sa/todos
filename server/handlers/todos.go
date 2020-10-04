package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) ListTodos(ctx *fiber.Ctx) error {
	todos, _ := h.Repo.ListTodos(ctx.Context())
	b := make([]interface{}, len(todos))
	for i := range todos {
		b[i] = todos[i]
	}
	return ctx.JSON(b)

}

func (h *Handlers) SetTodo(ctx *fiber.Ctx) error {
	type request struct {
		Name string `json name`
	}
	var body request
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})

	}
	if len(body.Name) <= 2 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "name not long enough",
		})
	}
	err = h.Repo.SetTodo(ctx.Context(), body.Name)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
	}

	return ctx.JSON(err)

}
