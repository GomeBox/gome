package mocks

import (
	"github.com/GomeBox/gome/adapters/audio"
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/adapters/input"
)

type System struct {
	OnInitialize      func() error
	OnUpdate          func() error
	OnInput           func() input.Adapters
	OnGraphics        func() graphics.Adapters
	OnAudio           func() audio.Adapters
	CallCntInitialize int
	CallCntUpdate     int
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

func (s *System) Input() input.Adapters {
	if s.OnInput != nil {
		return s.OnInput()
	}
	return nil
}

func (s *System) Graphics() graphics.Adapters {
	if s.OnGraphics != nil {
		return s.OnGraphics()
	}
	return nil
}

func (s *System) Audio() audio.Adapters {
	if s.OnAudio != nil {
		return s.OnAudio()
	}
	return nil
}
