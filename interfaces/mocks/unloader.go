package mocks

type Unloader struct {
	OnUnload   func() error
	OnUnloaded func() bool
}

func (u Unloader) Unload() error {
	if u.OnUnload != nil {
		return u.OnUnload()
	}
	return nil
}

func (u Unloader) Unloaded() bool {
	if u.OnUnloaded != nil {
		return u.OnUnloaded()
	}
	return false
}
