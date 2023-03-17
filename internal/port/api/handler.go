package api

import (
	"github.com/amirhnajafiz/DJaaS/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Repository repository.Repository
}

func (h *Handler) HandlerRegister() fiber.Router {
	app := fiber.New().Group("/api")

	app.Post("/submit/:question_id", h.MakeSubmit)
	app.Get("/submit/:question_id", h.GetQuestionSubmitsOfUser)
	app.Get("/submit/:question_id/all", h.GetQuestionSubmits)

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
