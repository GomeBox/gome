package gome

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/adapters/input"
	"github.com/GomeBox/gome/internal/core"
)

type Context interface {
	QuitGame()
	Graphics() graphics.Interface
	Input() input.Interface
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

func (context *contextWrapper) Graphics() graphics.Interface {
	return context.runner.AdapterSystem().Graphics()
}

func (context *contextWrapper) Input() input.Interface {
	return context.runner.AdapterSystem().Input()
}
