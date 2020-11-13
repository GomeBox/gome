package mocks

import (
	"github.com/GomeBox/gome/internal/input/interfaces"
)

type System struct {
	CallCntUpdate int
	OnKeyboard    func() interfaces.Keyboard
}

func (s *System) Keyboard() interfaces.Keyboard {
	if s.OnKeyboard != nil {
		return s.OnKeyboard()
	}
	return nil
}

func (s *System) Update() error {
	s.CallCntUpdate++
	return nil
}
