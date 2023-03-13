package cmd

import "github.com/spf13/cobra"

type HTTP struct{}

func (h HTTP) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "http",
		Short: "DJaaS http server",
		Run: func(_ *cobra.Command, _ []string) {
			h.main()
		},
	}
}

func (h HTTP) main() {

}
