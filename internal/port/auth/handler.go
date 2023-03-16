package auth

import (
	"github.com/amirhnajafiz/DJaaS/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Repository repository.Repository
}

func (h *Handler) Register(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) Login(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) Authenticate(ctx *fiber.Ctx) error {
	return nil
}
