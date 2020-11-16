package graphics

import (
	"github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/primitives"
)

type WindowSettings struct {
	Fullscreen bool
	Resolution primitives.Dimensions
	Title      string
}

func (w *WindowSettings) SetFullscreen(value bool) interfaces.WindowSettings {
	w.Fullscreen = value
	return w
}

func (w *WindowSettings) SetResolution(width, height int) interfaces.WindowSettings {
	w.Resolution.Width = width
	w.Resolution.Height = height
	return w
}

func (w *WindowSettings) SetTitle(value string) interfaces.WindowSettings {
	w.Title = value
	return w
}
