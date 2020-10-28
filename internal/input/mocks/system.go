package mocks

import "github.com/GomeBox/gome/input"

type System struct {
	CallCntUpdate int
}

func (s *System) Keyboard() input.Keyboard {
	panic("implement me")
}

func (s *System) Update() error {
	s.CallCntUpdate++
	return nil
}
