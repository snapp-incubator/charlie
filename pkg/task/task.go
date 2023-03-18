package task

type Task struct {
	Inputs   map[string]interface{}
	Function func(inputs map[string]interface{}) error
}

func (t Task) Execute() error {
	return t.Function(t.Inputs)
}
