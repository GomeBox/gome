package mocks

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/primitives"
)

type WindowAdapter struct {
}

func (w *WindowAdapter) OpenWindow(windowSettings *graphics.WindowSettings) error {
	panic("implement me")
}

func (w *WindowAdapter) IsOpen() bool {
	panic("implement me")
}

func (w *WindowAdapter) Size() primitives.Dimensions {
	panic("implement me")
}
