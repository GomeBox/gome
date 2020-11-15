package graphics

import (
	"github.com/GomeBox/gome/primitives"
)

type WindowMock struct {
	OnOpen func(settings WindowSettings) error
}

func (w *WindowMock) Dimensions() (primitives.Dimensions, error) {
	panic("implement me")
}

func (w *WindowMock) Open(settings WindowSettings) error {
	if w.OnOpen != nil {
		return w.OnOpen(settings)
	}
	return nil
}
