package processor

import (
	"github.com/amirhnajafiz/DJaaS/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type Processor struct {
	Repository repository.Repository
}

func (p *Processor) RegisterHandler() fiber.Router {
	app := fiber.New().Group("/processor")

	app.Put("/:submit_id", func(ctx *fiber.Ctx) error {
		return nil
	})

	return app
}