package mocks

import (
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/game"
)

type GameMock struct {
}

func (game *GameMock) CreateAdapters() (adapters.System, error) {
	return new(AdapterSystem), nil
}

func (game *GameMock) Initialize(context game.Context) error {
	return nil
}

func (game *GameMock) Update(timeDelta float32, context game.Context) (keepRunning bool, err error) {
	return true, nil
}

func (game *GameMock) Draw(timeDelta float32, context game.Context) error {
	return nil
}
