package mocks

import (
	gomeInterfaces "github.com/GomeBox/gome/interfaces"
	graphics "github.com/GomeBox/gome/internal/graphics/interfaces"
	input "github.com/GomeBox/gome/internal/input/interfaces"
)

type SystemsFactory struct {
	OnCreateGraphicsSystem      func() graphics.System
	OnCreateInputSystem         func() input.System
	OnCreateAudioSystem         func() gomeInterfaces.Audio
	CallCntCreateGraphicsSystem int
	CallCntCreateInputSystem    int
	CallCntCreateAudioSystem    int
}

func (s *SystemsFactory) CreateGraphicsSystem() graphics.System {
	s.CallCntCreateGraphicsSystem++
	if s.OnCreateGraphicsSystem != nil {
		return s.OnCreateGraphicsSystem()
	}
	return nil
}

func (s *SystemsFactory) CreateInputSystem() input.System {
	s.CallCntCreateInputSystem++
	if s.OnCreateInputSystem != nil {
		return s.OnCreateInputSystem()
	}
	return nil
}

func (s *SystemsFactory) CreateAudioSystem() gomeInterfaces.Audio {
	s.CallCntCreateAudioSystem++
	if s.OnCreateAudioSystem != nil {
		return s.OnCreateAudioSystem()
	}
	return nil
}
