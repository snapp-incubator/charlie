package service

import (
	"github.com/amirhnajafiz/DJaaS/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Repository *repository.Repository
}

func (h *Handler) HandlerRegister() *fiber.App {
	app := fiber.New()

	common := app.Group("/")
	admin := app.Group("/admin")

	admin.Get("/user", h.GetAllUsers)
	admin.Patch("/user", h.UpdateUser)
	admin.Get("/class", h.GetAllClasses)
	admin.Put("/class/:class_id/user/:user_id", h.AddUserToClass)
	admin.Delete("/class/:class_id/user/:user_id", h.RemoveUserFromClass)
	admin.Get("/question", h.GetAllQuestions)

	common.Put("/class", h.CreateClass)
	common.Delete("/class", h.DeleteClass)
	common.Get("/class", h.GetAllUserClasses)
	common.Get("/class/:class_id", h.GetSingleClass)
	common.Get("/class/:class_id/user", h.GetClassUsers)
	common.Get("/class/:class_id/question", h.GetClassQuestions)
	common.Put("/class/:class_id/question", h.CreateQuestion)
	common.Get("/class/:class_id/question/:question_id", h.GetSingleQuestion)
	common.Delete("/class/:class_id/question/:question_id", h.RemoveQuestion)

	return app
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
