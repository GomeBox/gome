package mocks

import (
	"github.com/GomeBox/gome"
	"github.com/GomeBox/gome/adapters"
)

type GameMock struct {
}

func (game *GameMock) Initialize() (adapters.System, error) {
	return new(AdapterSystem), nil
}

func (game *GameMock) Update(timeDelta float32, context gome.Context) error {
	context.QuitGame()
	return nil
}

func (game *GameMock) Draw(timeDelta float32, context gome.Context) error {
	return nil
}
