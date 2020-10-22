package graphics

import (
	"errors"
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/primitives"
)

type Window interface {
	Dimensions() (primitives.Dimensions, error)
}

func newWindow(adapter graphics.WindowAdapter) (Window, error) {
	if adapter == nil {
		return nil, errors.New("argument nil: adapter")
	}
	w := new(window)
	w.adapter = adapter
	return w, nil
}

type window struct {
	adapter graphics.WindowAdapter
}

func (w *window) Dimensions() (primitives.Dimensions, error) {
	if !w.adapter.IsOpen() {
		return primitives.Dimensions{}, errors.New("window is not open and has no dimensions")
	}
	return w.adapter.Size(), nil
}
