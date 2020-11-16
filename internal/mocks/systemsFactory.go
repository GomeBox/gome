package mocks

import (
	"github.com/GomeBox/gome/internal/game/audio"
	"github.com/GomeBox/gome/internal/game/graphics"
	"github.com/GomeBox/gome/internal/game/input"
)

type SystemsFactory struct {
	CallCntCreateAudioSystem    int
	CallCntCreateGraphicsSystem int
	CallCntCreateInputSystem    int
	OnCreateGraphicsSystem      func() graphics.System
}

func (s *SystemsFactory) CreateAudioSystem() audio.System {
	s.CallCntCreateAudioSystem++
	return nil
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
	return nil
}
