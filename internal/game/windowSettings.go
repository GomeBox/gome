package game

import (
	"github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/primitives"
)

type windowSettings struct {
	fullscreen bool
	resolution primitives.Dimensions
	title      string
}

func (w *windowSettings) Fullscreen(value bool) interfaces.WindowSettings {
	w.fullscreen = value
	return w
}

func (w *windowSettings) Resolution(width, height int) interfaces.WindowSettings {
	w.resolution.Height = height
	w.resolution.Width = width
	return w
}

func (w *windowSettings) Title(value string) interfaces.WindowSettings {
	w.title = value
	return w
}
