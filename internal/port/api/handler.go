package api

import (
	"github.com/amirhnajafiz/DJaaS/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Repository repository.Repository
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