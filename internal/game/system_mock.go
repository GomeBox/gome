package game

import (
	"github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/internal/game/graphics"
)

type systemMock struct {
	OnGraphics        func() graphics.System
	OnUpdate          func() error
	OnInitialize      func() error
	OnOpenGameWindow  func(settings graphics.WindowSettings) error
	OnContext         func() interfaces.Context
	CallCntUpdate     int
	CallCntInitialize int
}

func (s *systemMock) Context() interfaces.Context {
	if s.OnContext != nil {
		return s.OnContext()
	}
	return nil
}

func (s *systemMock) OpenGameWindow(settings graphics.WindowSettings) error {
	if s.OnOpenGameWindow != nil {
		return s.OnOpenGameWindow(settings)
	}
	return nil
}

func (s *systemMock) Initialize() error {
	s.CallCntInitialize++
	if s.OnInitialize != nil {
		return s.OnInitialize()
	}
	return nil
}

func (s *systemMock) Update() error {
	s.CallCntUpdate++
	if s.OnUpdate != nil {
		return s.OnUpdate()
	}
	return nil
}

func (s *systemMock) Graphics() graphics.System {
	if s.OnGraphics != nil {
		return s.OnGraphics()
	}
	return nil
}
