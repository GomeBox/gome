package gome

import (
	"github.com/GomeBox/gome/adapters/input"
	"github.com/GomeBox/gome/graphics"
	"github.com/GomeBox/gome/internal/core"
)

//Context is passed to Game.Update and Game.Draw to access the game's systems, etc.
type Context interface {
	//QuitGame stops the game loop and causes the Gome.Run to return
	QuitGame()
	//Graphics returns the graphics adapter
	Graphics() graphics.System
	//Graphics returns the input adapter
	Input() input.Port
}

type contextWrapper struct {
	runner   core.GameRunner
	graphics graphics.System
}

func newContextWrapper(runner core.GameRunner) *contextWrapper {
	context := new(contextWrapper)
	context.runner = runner
	textureLoader := runner.AdapterSystem().Graphics().TextureLoader()
	windowAdapter := runner.AdapterSystem().Graphics().WindowAdapter()
	fontLoader := runner.AdapterSystem().Graphics().FontLoader()
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

func (context *contextWrapper) Input() input.Port {
	return context.runner.AdapterSystem().Input()
}
