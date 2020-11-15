package input

type SystemMock struct {
	CallCntUpdate int
	OnKeyboard    func() Keyboard
}

func (s *SystemMock) Keyboard() Keyboard {
	if s.OnKeyboard != nil {
		return s.OnKeyboard()
	}
	return nil
}

func (s *SystemMock) Update() error {
	s.CallCntUpdate++
	return nil
}
