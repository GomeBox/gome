package mocks

type Player struct {
}

func (p Player) Unload() error {
	return nil
}

func (p Player) Play() error {
	return nil
}

func (p Player) Unloaded() bool {
	return false
}
