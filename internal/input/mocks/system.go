package mocks

import (
	"github.com/GomeBox/gome/internal/input/interfaces"
)

type System struct {
	CallCntUpdate int
}

func (s *System) Keyboard() interfaces.Keyboard {
	panic("implement me")
}

func (s *System) Update() error {
	s.CallCntUpdate++
	return nil
}
