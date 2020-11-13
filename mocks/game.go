package mocks

import (
	"github.com/GomeBox/gome"
	"github.com/GomeBox/gome/adapters"
)

type GameMock struct {
}

func (game *GameMock) CreateAdapters() (adapters.System, error) {
	return new(AdapterSystem), nil
}

func (game *GameMock) Initialize(context gome.Context) error {
	return nil
}

func (game *GameMock) Update(timeDelta float32, context gome.Context) (keepRunning bool, err error) {
	return true, nil
}

func (game *GameMock) Draw(timeDelta float32, context gome.Context) error {
	return nil
}
