package graphics

import (
	"errors"
	adapters "github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/internal/game/graphics"
	"github.com/GomeBox/gome/primitives"
)

type Window interface {
	Dimensions() (primitives.Dimensions, error)
	Open(settings graphics.WindowSettings) error
}

func newWindow(adapter adapters.WindowAdapter) Window {
	w := new(window)
	w.adapter = adapter
	return w
}

type window struct {
	adapter adapters.WindowAdapter
}

func (w *window) Dimensions() (primitives.Dimensions, error) {
	if !w.adapter.IsOpen() {
		return primitives.Dimensions{}, errors.New("window is not open and has no dimensions")
	}
	return w.adapter.Size(), nil
}

func (w *window) Open(settings graphics.WindowSettings) error {
	ws := adapters.WindowSettings{
		Fullscreen: settings.Fullscreen,
		Resolution: settings.Resolution,
		Title:      settings.Title,
	}
	return w.adapter.OpenWindow(&ws)
}
