package cmd

import (
	"github.com/spf13/cobra"
)

type Processor struct{}

func (p Processor) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "processor",
		Short: "DJaaS processor service",
		Run: func(_ *cobra.Command, _ []string) {
			p.main()
		},
	}
}

func (p Processor) main() {}
