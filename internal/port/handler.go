package port

import (
	"github.com/amirhnajafiz/DJaaS/internal/port/api"
	"github.com/amirhnajafiz/DJaaS/internal/port/auth"
	"github.com/amirhnajafiz/DJaaS/internal/port/service"
	"github.com/amirhnajafiz/DJaaS/pkg/enum"

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
	svc.Put("/class", h.Auth.CheckRole(enum.RoleAdmin, enum.RoleTeacher), h.Service.CreateClass)
	svc.Delete("/class/:class_id", h.Auth.CheckRole(enum.RoleAdmin, enum.RoleTeacher), h.Service.DeleteClass)
	svc.Get("/class/:class_id", h.Service.GetSingleClass)
	svc.Get("/class/:class_id/user", h.Service.GetClassUsers)
	svc.Put("/class/:class_id/user", h.Auth.CheckRole(enum.RoleAdmin, enum.RoleTeacher), h.Service.AddUserToClass)
	svc.Delete("/class/:class_id/user", h.Auth.CheckRole(enum.RoleAdmin, enum.RoleTeacher), h.Service.RemoveUserFromClass)
	svc.Get("/class/:class_id/question", h.Service.GetClassQuestions)
	svc.Put("/class/:class_id/question", h.Auth.CheckRole(enum.RoleAdmin, enum.RoleTeacher), h.Service.CreateQuestion)
	svc.Delete("/class/:class_id/question", h.Auth.CheckRole(enum.RoleAdmin, enum.RoleTeacher), h.Service.RemoveQuestion)
	svc.Get("/class/:class_id/question/:question_id", h.Service.GetSingleQuestion)

	adminGroup := svc.Use(h.Auth.CheckRole(enum.RoleAdmin))

	adminGroup.Get("/users", h.Service.GetAllUsers)
	adminGroup.Patch("/users", h.Service.UpdateUser)
	adminGroup.Get("/classes", h.Service.GetAllClasses)
	adminGroup.Get("/question", h.Service.GetAllQuestions)

	return app
}
