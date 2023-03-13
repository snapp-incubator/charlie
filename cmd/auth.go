package cmd

import (
	"github.com/spf13/cobra"
)

type Auth struct{}

func (a Auth) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "auth",
		Short: "DJaaS authentication service",
		Run: func(_ *cobra.Command, _ []string) {
			a.main()
		},
	}
}

func (a Auth) main() {}
