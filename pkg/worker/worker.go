package worker

import (
	"log"

	"github.com/amirhnajafiz/DJaaS/pkg/task"
)

type Master struct {
	Cfg     Config
	Channel chan task.Task
}

type slave struct {
	Channel chan task.Task
}

func New(cfg Config) Master {
	channel := make(chan task.Task)

	return Master{
		Cfg:     cfg,
		Channel: channel,
	}
}

func (m *Master) Register() {
	for i := 0; i < m.Cfg.Maximum; i++ {
		tmp := slave{
			Channel: m.Channel,
		}

		go tmp.work()
	}
}

func (m *Master) Push(t *task.Task) {
	m.Channel <- *t
}

func (s slave) work() {
	for {
		t := <-s.Channel

		if err := t.Execute(); err != nil {
			log.Println(err)
		}
	}
}
