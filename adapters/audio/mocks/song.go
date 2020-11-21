package mocks

type Song struct {
	OnPlay   func() error
	OnUnload func() error
}

func (s *Song) Play() error {
	if s.OnPlay != nil {
		return s.OnPlay()
	}
	return nil
}

func (s *Song) Unload() error {
	if s.OnUnload != nil {
		return s.OnUnload()
	}
	return nil
}
