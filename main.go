package main

import (
	"log"

	"github.com/amirhnajafiz/DJaaS/cmd"
	"github.com/amirhnajafiz/DJaaS/internal/storage"

	"github.com/spf13/cobra"
)

func main() {
	root := cobra.Command{}

	db, err := storage.NewConnection(storage.Config{})
	if err != nil {
		log.Println(err)

		return
	}

	root.AddCommand(
		cmd.Service{}.Command(),
		cmd.API{}.Command(),
		cmd.Auth{}.Command(),
		cmd.Migrate{}.Command(),
		cmd.Processor{}.Command(),
	)

	if err := root.Execute(); err != nil {
		panic(err)
	}
}
