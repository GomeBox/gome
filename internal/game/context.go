package game

import (
	"github.com/GomeBox/gome/audio"
	"github.com/GomeBox/gome/graphics"
	"github.com/GomeBox/gome/input"
	"github.com/GomeBox/gome/internal/game/interfaces"
)

type context struct {
	runner     interfaces.Runner
	gameSystem interfaces.System
}

func newContext(runner interfaces.Runner, system interfaces.System) *context {
	context := new(context)
	context.runner = runner
	context.gameSystem = system
	return context
}

func (context *context) QuitGame() {
	context.runner.Quit()
}

func (context *context) Graphics() graphics.System {
	return context.gameSystem.Graphics()
}

func (context *context) Input() input.System {
	return context.gameSystem.Input()
}

func (context *context) Audio() audio.System {
	return context.gameSystem.Audio()
}

//TODO: should be done in runner
func (context *context) Update() error {
	return context.gameSystem.Update()
}
