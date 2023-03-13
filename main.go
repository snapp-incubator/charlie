package main

import (
	"github.com/amirhnajafiz/DJaaS/cmd"

	"github.com/spf13/cobra"
)

func main() {
	root := cobra.Command{}

	root.AddCommand(
		cmd.HTTP{}.Command(),
		cmd.API{}.Command(),
		cmd.Auth{}.Command(),
		cmd.Migrate{}.Command(),
		cmd.Processor{}.Command(),
	)

	if err := root.Execute(); err != nil {
		panic(err)
	}
}
