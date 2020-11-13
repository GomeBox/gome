package mocks

import (
	"github.com/GomeBox/gome/interfaces"
	graphics "github.com/GomeBox/gome/internal/graphics/interfaces"
	input "github.com/GomeBox/gome/internal/input/interfaces"
)

type System struct {
	OnGraphics        func() graphics.System
	OnInput           func() input.System
	OnAudio           func() interfaces.Audio
	OnUpdate          func() error
	OnInitialize      func() error
	CallCntUpdate     int
	CallCntInitialize int
}

func (s *System) Initialize() error {
	s.CallCntInitialize++
	if s.OnInitialize != nil {
		return s.OnInitialize()
	}
	return nil
}

func (s *System) Update() error {
	s.CallCntUpdate++
	if s.OnUpdate != nil {
		return s.OnUpdate()
	}
	return nil
}

func (s *System) Graphics() graphics.System {
	if s.OnGraphics != nil {
		return s.OnGraphics()
	}
	return nil
}

func (s *System) Input() input.System {
	if s.OnInput != nil {
		return s.OnInput()
	}
	return nil
}

func (s *System) Audio() interfaces.Audio {
	if s.OnAudio != nil {
		return s.OnAudio()
	}
	return nil
}
