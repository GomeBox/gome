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
	game *core.Game
}

func newContextWrapper(g *core.Game) *contextWrapper {
	context := new(contextWrapper)
	context.game = g
	return context
}

func (c *contextWrapper) QuitGame() {
	c.game.Quit()
}

func (c *contextWrapper) Graphics() graphics.Interface {
	return c.game.AdapterSystem().Graphics()
}

func (c *contextWrapper) Input() input.Interface {
	return c.game.AdapterSystem().Input()
}
