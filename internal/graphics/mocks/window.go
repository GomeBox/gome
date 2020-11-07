package mocks

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/primitives"
)

type Window struct {
	OnOpen func(settings *graphics.WindowSettings) error
}

func (w *Window) Dimensions() (primitives.Dimensions, error) {
	panic("implement me")
}

func (w *Window) Open(settings *graphics.WindowSettings) error {
	if w.OnOpen != nil {
		return w.OnOpen(settings)
	}
	return nil
}
