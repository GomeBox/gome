package input

import "github.com/GomeBox/gome/interfaces"

type SystemMock struct {
	CallCntUpdate int
	OnKeyboard    func() interfaces.Keyboard
}

func (s *SystemMock) Keyboard() interfaces.Keyboard {
	if s.OnKeyboard != nil {
		return s.OnKeyboard()
	}
	return nil
}

func (s *SystemMock) Update() error {
	s.CallCntUpdate++
	return nil
}
