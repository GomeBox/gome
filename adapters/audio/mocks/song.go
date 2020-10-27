package mocks

type Song struct {
	OnPlay func() error
}

func (s *Song) Play() error {
	if s.OnPlay != nil {
		return s.OnPlay()
	}
	return nil
}
