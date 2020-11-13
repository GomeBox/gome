package mocks

type Runner struct {
	OnQuit func()
}

func (r Runner) Run() error {
	panic("implement me")
}
