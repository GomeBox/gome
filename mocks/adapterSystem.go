package mocks

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/adapters/input"
)

type AdapterSystem struct {
}

func (adapterSystem *AdapterSystem) Initialize() error {
	return nil
}

func (adapterSystem *AdapterSystem) Update() error {
	return nil
}

func (adapterSystem *AdapterSystem) Input() input.Adapters {
	return new(InputAdapters)
}

func (adapterSystem *AdapterSystem) Graphics() graphics.Adapters {
	return new(GraphicsAdapters)
}
