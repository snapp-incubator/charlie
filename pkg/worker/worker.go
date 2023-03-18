package worker

import (
	"github.com/amirhnajafiz/DJaaS/pkg/task"
)

type Master struct {
	Cfg     Config
	Channel chan task.Task
}

type Slave struct {
	Channel chan task.Task
}

func New(cfg Config) Master {
	channel := make(chan task.Task)

	return Master{
		Cfg:     cfg,
		Channel: channel,
	}
}
