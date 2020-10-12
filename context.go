package game

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/adapters/input"
)

type Context interface {
	QuitGame()
	Graphics() graphics.Interface
	Input() input.Interface
}

func newContext(g *gome) Context {
	context := new(context)
	context.gome = g
	return context
}

type context struct {
	gome *gome
}

func (c *context) QuitGame() {
	c.gome.Quit()
}

func (c *context) Graphics() graphics.Interface {
	return c.gome.adapterSystem.Graphics()
}

func (c *context) Input() input.Interface {
	return c.gome.adapterSystem.Input()
}
