package port

import (
	"github.com/amirhnajafiz/DJaaS/internal/port/api"
	"github.com/amirhnajafiz/DJaaS/internal/port/auth"
	"github.com/amirhnajafiz/DJaaS/internal/port/service"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	API     api.Handler
	Auth    auth.Handler
	Service service.Handler
}

func (h *Handler) Register() *fiber.App {
	app := fiber.New()

	authentication := app.Group("/auth")

	apiGroup := app.Group("/api").Use(h.Auth.Authenticate)
	svc := app.Group("/svc").Use(h.Auth.Authenticate)

	authentication.Put("/register", h.Auth.Register)
	authentication.Post("/login", h.Auth.Login)

	apiGroup.Post("/submit", h.API.MakeSubmit)
	apiGroup.Get("/submit/:question_id", h.API.GetQuestionSubmits)
	apiGroup.Get("/submit/:question_id/user", h.API.GetQuestionSubmitsOfUser)

	svc.Get("/class", h.Service.GetAllUserClasses)
	svc.Put("/class", h.Service.CreateClass)
	svc.Delete("/class/:class_id", h.Service.DeleteClass)
	svc.Get("/class/:class_id", h.Service.GetSingleClass)
	svc.Get("/class/:class_id/user", h.Service.GetClassUsers)
	svc.Put("/class/:class_id/user", h.Service.AddUserToClass)
	svc.Delete("/class/:class_id/user", h.Service.RemoveUserFromClass)
	svc.Get("/class/:class_id/question", h.Service.GetClassQuestions)
	svc.Put("/class/:class_id/question", h.Service.CreateQuestion)
	svc.Delete("/class/:class_id/question", h.Service.RemoveQuestion)
	svc.Get("/class/:class_id/question/:question_id", h.Service.GetSingleQuestion)

	svc.Get("/users", h.Service.GetAllUsers)
	svc.Patch("/users", h.Service.UpdateUser)
	svc.Get("/classes", h.Service.GetAllClasses)
	svc.Get("/question", h.Service.GetAllQuestions)
}
