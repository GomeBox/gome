package gome

import (
	"github.com/GomeBox/gome/audio"
	"github.com/GomeBox/gome/graphics"
	"github.com/GomeBox/gome/input"
	"github.com/GomeBox/gome/internal/game/interfaces"
)

//Context is passed to Interface.Update and Interface.Draw to access the game's systems, etc.
type Context interface {
	//Graphics returns the graphics adapter
	Graphics() graphics.System
	//Graphics returns the input adapter
	Input() input.System
	//Audio returns the audio adapter
	Audio() audio.System
}

func newContext(gameSystem interfaces.System) Context {
	c := new(context)
	c.graphics = graphics.NewSystem(gameSystem.Graphics())
	c.input = input.NewSystem(gameSystem.Input())
	c.audio = audio.NewSystem(gameSystem.Audio())
	return c
}

type context struct {
	graphics graphics.System
	input    input.System
	audio    audio.System
}

func (context *context) Graphics() graphics.System {
	return context.graphics
}

func (context *context) Input() input.System {
	return context.input
}

func (context *context) Audio() audio.System {
	return context.audio
}
