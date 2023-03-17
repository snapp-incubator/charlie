package main

import (
	"log"

	"github.com/amirhnajafiz/DJaaS/cmd"
	"github.com/amirhnajafiz/DJaaS/internal/port/api"
	"github.com/amirhnajafiz/DJaaS/internal/port/auth"
	"github.com/amirhnajafiz/DJaaS/internal/port/processor"
	"github.com/amirhnajafiz/DJaaS/internal/port/service"
	"github.com/amirhnajafiz/DJaaS/internal/repository"
	"github.com/amirhnajafiz/DJaaS/internal/storage"

	"github.com/spf13/cobra"
)

func main() {
	root := cobra.Command{}

	// open database connection
	db, err := storage.NewConnection(storage.Config{})
	if err != nil {
		log.Println(err)

		return
	}

	// create repository
	r := repository.New(db)

	root.AddCommand(
		cmd.Service{
			Handler: service.Handler{
				Repository: r,
			},
		}.Command(),
		cmd.API{
			Handler: api.Handler{
				Repository: r,
			},
		}.Command(),
		cmd.Auth{
			Handler: auth.Handler{
				Repository: r,
			},
		}.Command(),
		cmd.Migrate{}.Command(),
		cmd.Processor{
			Handler: processor.Processor{
				Repository: r,
			},
		}.Command(),
	)

	if err := root.Execute(); err != nil {
		panic(err)
	}
}
