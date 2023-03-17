package cmd

import (
	"github.com/amirhnajafiz/DJaaS/internal/port"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

type API struct {
	Handlers port.Handler
}

func (a API) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "DJaaS api server",
		Run: func(_ *cobra.Command, _ []string) {
			a.main()
		},
	}
}

func (a API) main() {
	app := fiber.New()

	api := app.Group("/api")

	api.Post("/submit", a.Handlers.API.MakeSubmit)
	api.Get("/submit/:question_id", a.Handlers.API.GetQuestionSubmits)
	api.Get("/submit/:question_id/user", a.Handlers.API.GetQuestionSubmitsOfUser)
}
