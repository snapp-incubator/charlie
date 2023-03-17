package cmd

import (
	"log"

	"github.com/amirhnajafiz/DJaaS/internal/port/processor"

	"github.com/spf13/cobra"
)

type Processor struct {
	Handler processor.Processor
}

func (p Processor) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "processor",
		Short: "DJaaS processor service",
		Run: func(_ *cobra.Command, _ []string) {
			p.main()
		},
	}
}

func (p Processor) main() {
	app := p.Handler.RegisterHandler()

	go func() {
		if err := app.Listen(":8080"); err != nil {
			log.Println(err)
		}
	}()

	p.Handler.Process()
}
