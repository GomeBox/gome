package mocks

type Sound struct {
	OnPlay func() error
}

func (s Sound) Play() error {
	if s.OnPlay != nil {
		return s.OnPlay()
	}
	return nil
}
