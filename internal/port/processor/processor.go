package processor

import (
	"log"
	"strconv"

	"github.com/amirhnajafiz/DJaaS/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type Processor struct {
	Channel    chan uint
	Repository repository.Repository
}

func (p *Processor) Process() {
	for {
		id := <-p.Channel

		log.Println(id)
	}
}

func (p *Processor) RegisterHandler() fiber.Router {
	app := fiber.New().Group("/processor")

	app.Put("/:submit_id", func(ctx *fiber.Ctx) error {
		id, err := strconv.Atoi(ctx.Params("submit_id"))
		if err != nil {
			return fiber.ErrBadRequest
		}

		p.Channel <- uint(id)

		return ctx.SendStatus(fiber.StatusOK)
	})

	return app
}
