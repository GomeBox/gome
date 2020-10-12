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

func newContext(g *game) Context {
	context := new(context)
	context.game = g
	return context
}

type context struct {
	game *game
}

func (c *context) QuitGame() {
	c.game.Quit()
}

func (c *context) Graphics() graphics.Interface {
	return c.game.adapterSystem.Graphics()
}

func (c *context) Input() input.Interface {
	return c.game.adapterSystem.Input()
}
