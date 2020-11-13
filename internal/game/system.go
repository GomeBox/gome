package game

import (
	"github.com/GomeBox/gome/adapters"
	gomeInterfaces "github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/internal/game/interfaces"
	graphics "github.com/GomeBox/gome/internal/graphics/interfaces"
	input "github.com/GomeBox/gome/internal/input/interfaces"
)

func newSystem(adapters adapters.System, factory interfaces.SystemsFactory) interfaces.System {
	return &systemImpl{
		adapterSystem: adapters,
		graphics:      factory.CreateGraphicsSystem(),
		input:         factory.CreateInputSystem(),
		audio:         factory.CreateAudioSystem(),
	}
}

type systemImpl struct {
	adapterSystem adapters.System
	graphics      graphics.System
	input         input.System
	audio         gomeInterfaces.Audio
}

func (system *systemImpl) Initialize() error {
	return system.adapterSystem.Initialize()
}

func (system *systemImpl) Update() error {
	err := system.adapterSystem.Update()
	if err != nil {
		return err
	}
	return system.input.Update()
}

func (system *systemImpl) Input() input.System {
	return system.input
}

func (system *systemImpl) Audio() gomeInterfaces.Audio {
	return system.audio
}

func (system *systemImpl) Graphics() graphics.System {
	return system.graphics
}
