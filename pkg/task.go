package jolt

type Task struct {
	Environment
	Jobs []Job
}

func NewTask(name string) *Task {
	t := &Task{}
	t.reset()
	t.Name = name
	return t
}
