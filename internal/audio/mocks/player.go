package mocks

type Player struct {
	OnPlay      func() error
	CallCntPlay int
}

func (s *Player) Play() error {
	s.CallCntPlay++
	if s.OnPlay != nil {
		return s.OnPlay()
	}
	return nil
}
