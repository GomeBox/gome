package gome

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/adapters/input"
	"github.com/GomeBox/gome/internal/core"
)

//Context is passed to Game.Update and Game.Draw to access the game's systems, etc.
type Context interface {
	//QuitGame stops the game loop and causes the Gome.Run to return
	QuitGame()
	//Graphics returns the graphics adapter
	Graphics() graphics.Port
	//Graphics returns the input adapter
	Input() input.Port
}

type contextWrapper struct {
	runner core.GameRunner
}

func newContextWrapper(runner core.GameRunner) *contextWrapper {
	context := new(contextWrapper)
	context.runner = runner
	return context
}

func (context *contextWrapper) QuitGame() {
	context.runner.Quit()
}

func (context *contextWrapper) Graphics() graphics.Port {
	return context.runner.AdapterSystem().Graphics()
}

func (context *contextWrapper) Input() input.Port {
	return context.runner.AdapterSystem().Input()
}
