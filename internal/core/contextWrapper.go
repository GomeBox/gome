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
}

func newContextWrapper(system adapters.System, runner GameRunner) *contextWrapper {
	context := new(contextWrapper)
	context.runner = runner
	textureLoader := system.Graphics().TextureLoader()
	windowAdapter := system.Graphics().WindowAdapter()
	fontLoader := system.Graphics().FontLoader()
	context.graphics = graphics.NewSystem(
		graphics.Adapters{
			TextureLoader: textureLoader,
			WindowAdapter: windowAdapter,
			FontLoader:    fontLoader})

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
