package game

import (
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/internal/game/audio"
	"github.com/GomeBox/gome/internal/game/graphics"
	"github.com/GomeBox/gome/internal/game/input"
)

type System interface {
	Initialize() error
	Update() error
	Graphics() graphics.System
	Context() interfaces.Context
	OpenGameWindow(settings graphics.WindowSettings) error
}

func NewSystem(adapters adapters.System, graphics graphics.System, audio audio.System, input input.System) System {
	return &system{
		adapterSystem: adapters,
		graphics:      graphics,
		input:         input,
		audio:         audio,
	}
}

type system struct {
	adapterSystem adapters.System
	graphics      graphics.System
	input         input.System
	audio         audio.System
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

func (system *system) Graphics() graphics.System {
	return system.graphics
}

func (system *system) Context() interfaces.Context {
	return nil
}

func (system *system) OpenGameWindow(settings graphics.WindowSettings) error {
	return system.graphics.OpenWindow(settings)
}
