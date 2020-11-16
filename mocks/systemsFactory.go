package mocks

import (
	audioAdapters "github.com/GomeBox/gome/adapters/audio"
	graphicsAdapters "github.com/GomeBox/gome/adapters/graphics"
	inputAdapters "github.com/GomeBox/gome/adapters/input"
	"github.com/GomeBox/gome/internal/game/audio"
	"github.com/GomeBox/gome/internal/game/graphics"
	"github.com/GomeBox/gome/internal/game/input"
)

type SystemsFactory struct {
	CallCntCreateAudioSystem    int
	CallCntCreateGraphicsSystem int
	CallCntCreateInputSystem    int
	OnCreateGraphicsSystem      func(adapters graphicsAdapters.Adapters) graphics.System
}

func (s *SystemsFactory) CreateAudioSystem(adapters audioAdapters.Adapters) audio.System {
	s.CallCntCreateAudioSystem++
	return nil
}

func (s *SystemsFactory) CreateGraphicsSystem(adapters graphicsAdapters.Adapters) graphics.System {
	s.CallCntCreateGraphicsSystem++
	if s.OnCreateGraphicsSystem != nil {
		return s.OnCreateGraphicsSystem(adapters)
	}
	return nil
}

func (s *SystemsFactory) CreateInputSystem(adapters inputAdapters.Adapters) input.System {
	s.CallCntCreateInputSystem++
	return nil
}
