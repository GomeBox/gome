package core

import (
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/graphics"
	"github.com/GomeBox/gome/input"
	internalGraphics "github.com/GomeBox/gome/internal/graphics"
	internalInput "github.com/GomeBox/gome/internal/input"
)

type contextWrapper struct {
	runner   GameRunner
	graphics graphics.System
	input    input.System
	adapters adapters.System
}

func newContextWrapper(system adapters.System, runner GameRunner) (*contextWrapper, error) {
	context := new(contextWrapper)
	context.runner = runner
	context.adapters = system
	textureLoader := system.Graphics().TextureLoader()
	textureCreator := system.Graphics().TextureCreator()
	windowAdapter := system.Graphics().WindowAdapter()
	fontLoader := system.Graphics().FontLoader()
	tmpGraphics, err := internalGraphics.NewSystem(
		internalGraphics.Adapters{
			TextureCreator: textureCreator,
			TextureLoader:  textureLoader,
			WindowAdapter:  windowAdapter,
			FontLoader:     fontLoader})
	if err != nil {
		return nil, err
	}
	context.graphics = tmpGraphics
	context.input = internalInput.NewSystem(
		internalInput.Adapters{
			Keyboard: system.Input().Keyboard()})

	return context, nil
}

func (context *contextWrapper) QuitGame() {
	context.runner.Quit()
}

func (context *contextWrapper) Graphics() graphics.System {
	return context.graphics
}

func (context *contextWrapper) Input() input.System {
	return context.input
}

func (context *contextWrapper) Update() error {
	err := context.adapters.Update()
	if err != nil {
		return err
	}
	err = context.Input().Update()
	if err != nil {
		return err
	}
	return err
}
