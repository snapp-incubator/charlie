package cmd

import (
	"log"

	"github.com/amirhnajafiz/DJaaS/internal/port/api"

	"github.com/spf13/cobra"
)

type API struct {
	Handler api.Handler
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
	app := a.Handler.HandlerRegister()

	if err := app.Listen(":8080"); err != nil {
		log.Println(err)
	}
}
