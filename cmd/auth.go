package cmd

import (
	"log"

	"github.com/amirhnajafiz/DJaaS/internal/port/auth"

	"github.com/spf13/cobra"
)

type Auth struct {
	Handler auth.Handler
}

func (a Auth) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "auth",
		Short: "DJaaS authentication service",
		Run: func(_ *cobra.Command, _ []string) {
			a.main()
		},
	}
}

func (a Auth) main() {
	app := a.Handler.HandlerRegister()

	if err := app.Listen(":8080"); err != nil {
		log.Println(err)
	}
}
