package gome

import (
	audioAdapters "github.com/GomeBox/gome/adapters/audio"
	graphicsAdapters "github.com/GomeBox/gome/adapters/graphics"
	inputAdapters "github.com/GomeBox/gome/adapters/input"
	"github.com/GomeBox/gome/internal/audio"
	gameAudio "github.com/GomeBox/gome/internal/game/audio"
	gameGraphics "github.com/GomeBox/gome/internal/game/graphics"
	gameInput "github.com/GomeBox/gome/internal/game/input"
	"github.com/GomeBox/gome/internal/graphics"
	"github.com/GomeBox/gome/internal/input"
)

type SystemsFactory interface {
	CreateAudioSystem(adapters audioAdapters.Adapters) gameAudio.System
	CreateGraphicsSystem(adapters graphicsAdapters.Adapters) gameGraphics.System
	CreateInputSystem(adapters inputAdapters.Adapters) gameInput.System
}

type systemsFactory struct {
}

func (s *systemsFactory) CreateAudioSystem(adapters audioAdapters.Adapters) gameAudio.System {
	return audio.NewSystem(adapters)
}

func (s *systemsFactory) CreateGraphicsSystem(adapters graphicsAdapters.Adapters) gameGraphics.System {
	return graphics.NewSystem(adapters)
}

func (s *systemsFactory) CreateInputSystem(adapters inputAdapters.Adapters) gameInput.System {
	return input.NewSystem(adapters)
}
