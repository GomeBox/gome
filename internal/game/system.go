package game

import (
	"github.com/GomeBox/gome/adapters"
	gomeInterfaces "github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/internal/audio"
	"github.com/GomeBox/gome/internal/game/interfaces"
	"github.com/GomeBox/gome/internal/graphics"
	input "github.com/GomeBox/gome/internal/input/interfaces"
)

func NewSystem(adapters adapters.System, factory interfaces.SystemsFactory) interfaces.System {
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
	audio         audio.System
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

func (system *systemImpl) Graphics() graphics.System {
	return system.graphics
}

func (system *systemImpl) Context() gomeInterfaces.Context {
	return nil
}

func (system *systemImpl) OpenGameWindow(settings graphics.WindowSettings) error {
	return system.graphics.Window().Open(settings)
}
