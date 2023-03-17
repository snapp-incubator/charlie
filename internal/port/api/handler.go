package api

import (
	"github.com/amirhnajafiz/DJaaS/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Repository *repository.Repository
}

func (h *Handler) HandlerRegister() *fiber.App {
	app := fiber.New()

	api := app.Group("/api")

	api.Post("/submit/:question_id", h.MakeSubmit)
	api.Get("/submit/:question_id", h.GetQuestionSubmitsOfUser)
	api.Get("/submit/:question_id/all", h.GetQuestionSubmits)

	return app
}

func (h *Handler) MakeSubmit(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) GetQuestionSubmits(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) GetQuestionSubmitsOfUser(ctx *fiber.Ctx) error {
	return nil
}
