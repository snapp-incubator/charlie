package cmd

import (
	"log"

	"github.com/amirhnajafiz/DJaaS/internal/port/service"

	"github.com/spf13/cobra"
)

type Service struct {
	Handler service.Handler
}

func (s Service) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "http",
		Short: "DJaaS http server",
		Run: func(_ *cobra.Command, _ []string) {
			s.main()
		},
	}
}

func (s Service) main() {
	app := s.Handler.HandlerRegister()

	if err := app.Listen(":8080"); err != nil {
		log.Println(err)
	}
}
