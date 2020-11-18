package mocks

type Sound struct {
	OnPlay   func() error
	OnUnload func() error
}

func (s *Sound) Play() error {
	if s.OnPlay != nil {
		return s.OnPlay()
	}
	return nil
}

func (s *Sound) Unload() error {
	if s.OnUnload != nil {
		return s.OnUnload()
	}
	return nil
}
