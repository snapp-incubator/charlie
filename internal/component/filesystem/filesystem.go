package filesystem

import (
	"os"
)

type System struct {
	Cfg Config
}

func New(cfg Config) *System {
	return &System{
		Cfg: cfg,
	}
}

func (s *System) GetFile(name string) (os.File, error) {
	return os.File{}, nil
}
