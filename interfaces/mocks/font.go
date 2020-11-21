package mocks

import (
	"github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/primitives"
)

type Font struct {
}

func (f Font) Unload() error {
	panic("implement me")
}

func (f Font) Unloaded() bool {
	panic("implement me")
}

func (f Font) CreateText(value string, color primitives.Color) (interfaces.Text, error) {
	panic("implement me")
}
