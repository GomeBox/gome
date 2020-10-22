package internal

import (
	"errors"
	"github.com/GomeBox/gome/game"
	internalGame "github.com/GomeBox/gome/internal/game"
)

type GomeImpl struct {
	GameRunner   internalGame.GameRunner
	GameSettings game.Settings
}

//Run starts a game and starts the game loop. Returns, when the game is ended
func (gome *GomeImpl) Run(game game.Interface) error {
	if gome.GameRunner.Running() {
		return errors.New("game is already running")
	}
	err := gome.GameRunner.Init(game.CreateAdapters, gome.GameSettings)
	if err != nil {
		return err
	}
	err = gome.GameRunner.Loop(game.Initialize, game.Update, game.Draw)
	if err != nil {
		return err
	}
	return nil
}

func (gome *GomeImpl) Settings() game.Settings {
	return gome.GameSettings
}
