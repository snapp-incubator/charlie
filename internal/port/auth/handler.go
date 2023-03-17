package auth

import (
	"github.com/amirhnajafiz/DJaaS/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Repository *repository.Repository
}

func (h *Handler) HandlerRegister() *fiber.App {
	app := fiber.New()

	auth := app.Group("/auth")

	auth.Put("/user", h.Register)
	auth.Post("/user", h.Login)

	return app
}

func (h *Handler) Register(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) Login(ctx *fiber.Ctx) error {
	return nil
}
