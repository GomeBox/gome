package mocks

import (
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/game"
)

type GameMock struct {
}

func (game *GameMock) Initialize() (adapters.System, error) {
	return new(AdapterSystem), nil
}

func (game *GameMock) Update(timeDelta float32, context game.Context) error {
	context.QuitGame()
	return nil
}

func (game *GameMock) Draw(timeDelta float32, context game.Context) error {
	return nil
}
