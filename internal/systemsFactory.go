package internal

import (
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/internal/audio"
	gameAudio "github.com/GomeBox/gome/internal/game/audio"
	gameGraphics "github.com/GomeBox/gome/internal/game/graphics"
	gameInput "github.com/GomeBox/gome/internal/game/input"
	"github.com/GomeBox/gome/internal/graphics"
	"github.com/GomeBox/gome/internal/input"
)

type SystemsFactory interface {
	CreateAudioSystem() gameAudio.System
	CreateGraphicsSystem() gameGraphics.System
	CreateInputSystem() gameInput.System
}

func NewSystemsFactory(adapterSystem adapters.System) SystemsFactory {
	return &systemsFactory{adapterSystem: adapterSystem}
}

type systemsFactory struct {
	adapterSystem adapters.System
}

func (s *systemsFactory) CreateAudioSystem() gameAudio.System {
	return audio.NewSystem(s.adapterSystem.Audio())
}

func (s *systemsFactory) CreateGraphicsSystem() gameGraphics.System {
	return graphics.NewSystem(s.adapterSystem.Graphics())
}

func (s *systemsFactory) CreateInputSystem() gameInput.System {
	return input.NewSystem(s.adapterSystem.Input())
}
