package core

import (
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/graphics"
	"github.com/GomeBox/gome/input"
)

type contextWrapper struct {
	runner   GameRunner
	graphics graphics.System
	input    input.System
	adapters adapters.System
}

func newContextWrapper(system adapters.System, runner GameRunner) *contextWrapper {
	context := new(contextWrapper)
	context.runner = runner
	context.adapters = system
	textureLoader := system.Graphics().TextureLoader()
	textureCreator := system.Graphics().TextureCreator()
	windowAdapter := system.Graphics().WindowAdapter()
	fontLoader := system.Graphics().FontLoader()
	context.graphics = graphics.NewSystem(
		graphics.Adapters{
			TextureCreator: textureCreator,
			TextureLoader:  textureLoader,
			WindowAdapter:  windowAdapter,
			FontLoader:     fontLoader})
	context.input = input.NewSystem(
		input.Adapters{
			Keyboard: system.Input().Keyboard()})

	return context
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
