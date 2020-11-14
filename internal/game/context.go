package game

import (
	gomeInterfaces "github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/internal/game/interfaces"
)

func newContext(gameSystem interfaces.System) gomeInterfaces.Context {
	c := new(context)
	c.graphics = gameSystem.Graphics()
	c.input = gomeInterfaces.NewSystem(gameSystem.Input())
	c.audio = gameSystem.Audio()
	return c
}

type context struct {
	graphics gomeInterfaces.Graphics
	input    gomeInterfaces.Input
	audio    gomeInterfaces.Audio
}

func (context *context) Graphics() gomeInterfaces.Graphics {
	return context.graphics
}

func (context *context) Input() gomeInterfaces.Input {
	return context.input
}

func (context *context) Audio() gomeInterfaces.Audio {
	return context.audio
}
