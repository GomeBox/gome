package mocks

import (
	"github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/internal/game/graphics"
)

type System struct {
	OnGraphics        func() graphics.System
	OnUpdate          func() error
	OnInitialize      func() error
	OnOpenGameWindow  func(settings graphics.WindowSettings) error
	OnContext         func() interfaces.Context
	CallCntUpdate     int
	CallCntInitialize int
}

func (s *System) Context() interfaces.Context {
	if s.OnContext != nil {
		return s.OnContext()
	}
	return nil
}

func (s *System) OpenGameWindow(settings graphics.WindowSettings) error {
	if s.OnOpenGameWindow != nil {
		return s.OnOpenGameWindow(settings)
	}
	return nil
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
