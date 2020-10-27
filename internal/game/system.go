package game

import (
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/internal/audio"
	"github.com/GomeBox/gome/internal/game/interfaces"
	"github.com/GomeBox/gome/internal/graphics"
	"github.com/GomeBox/gome/internal/input"
)

func NewSystem(adapters adapters.System) interfaces.System {
	factory := newSystemsFactory(adapters)
	return &system{
		adapterSystem: adapters,
		graphics:      factory.CreateGraphicsSystem(),
		input:         factory.CreateInputSystem(),
		audio:         factory.CreateAudioSystem(),
	}
}

type system struct {
	adapterSystem adapters.System
	graphics      *graphics.System
	input         *input.System
	audio         *audio.System
}

func (system *system) Initialize() error {
	return system.adapterSystem.Initialize()
}

func (system *system) Update() error {
	err := system.adapterSystem.Update()
	if err != nil {
		return err
	}
	return system.input.Update()
}

func (system *system) Input() *input.System {
	return system.input
}

func (system *system) Audio() *audio.System {
	return system.audio
}

func (system *system) Graphics() *graphics.System {
	return system.graphics
}
