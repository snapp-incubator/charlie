package cmd

import (
	"github.com/spf13/cobra"
)

type Service struct{}

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

}
