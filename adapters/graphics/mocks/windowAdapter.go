package mocks

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/primitives"
)

type WindowAdapter struct {
	OnOpenWindow func(windowSettings *graphics.WindowSettings) error
	OnIsOpen     func() bool
	OnSize       func() primitives.Dimensions
}

func (w *WindowAdapter) OpenWindow(windowSettings *graphics.WindowSettings) error {
	if w.OnOpenWindow != nil {
		return w.OnOpenWindow(windowSettings)
	}
	return nil
}

func (w *WindowAdapter) IsOpen() bool {
	if w.OnIsOpen != nil {
		return w.OnIsOpen()
	}
	return true
}

func (w *WindowAdapter) Size() primitives.Dimensions {
	if w.OnSize != nil {
		return w.OnSize()
	}
	return primitives.Dimensions{}
}
