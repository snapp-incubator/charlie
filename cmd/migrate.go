package cmd

import (
	"log"

	"github.com/amirhnajafiz/DJaaS/internal/model"
	"github.com/amirhnajafiz/DJaaS/internal/storage"

	"github.com/spf13/cobra"
)

type Migrate struct {
	Cfg storage.Config
}

func (m Migrate) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "DJaaS migrate database",
		Run: func(_ *cobra.Command, _ []string) {
			m.main()
		},
	}
}

func (m Migrate) main() {
	db, err := storage.NewConnection(m.Cfg)
	if err != nil {
		log.Println(err)

		return
	}

	if er := db.AutoMigrate(
		&model.Class{},
		&model.User{},
		&model.Question{},
		&model.Submit{},
	); er != nil {
		log.Println(er)

		return
	}

	log.Println("Migration succeed")
}
