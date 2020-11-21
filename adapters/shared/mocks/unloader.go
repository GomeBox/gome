package mocks

type Unloader struct {
	OnUnload func() error
}

func (u Unloader) Unload() error {
	if u.OnUnload != nil {
		return u.OnUnload()
	}
	return nil
}
