package game

import (
	gomeInterfaces "github.com/GomeBox/gome/interfaces"
)

func newContext(
	graphics gomeInterfaces.Graphics,
	input gomeInterfaces.Input,
	audio gomeInterfaces.Audio) gomeInterfaces.Context {
	c := new(context)
	c.graphics = graphics
	c.input = input
	c.audio = audio
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
