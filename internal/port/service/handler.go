package service

import (
	"github.com/amirhnajafiz/DJaaS/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Repository repository.Repository
}

func (h *Handler) GetAllUsers(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) UpdateUser(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) CreateClass(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) DeleteClass(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) GetSingleClass(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) GetAllUserClasses(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) GetAllClasses(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) GetClassUsers(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) GetClassQuestions(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) AddUserToClass(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) RemoveUserFromClass(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) CreateQuestion(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) RemoveQuestion(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) GetSingleQuestion(ctx *fiber.Ctx) error {
	return nil
}

func (h *Handler) GetAllQuestions(ctx *fiber.Ctx) error {
	return nil
}
