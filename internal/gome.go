package internal

import (
	"errors"
	"github.com/GomeBox/gome/game"
	"github.com/GomeBox/gome/internal/core"
)

type GomeImpl struct {
	GameRunner core.GameRunner
}

//Run starts a game and starts the game loop. Returns, when the game is ended
func (gome *GomeImpl) Run(game game.Interface, settings game.Settings) error {
	if gome.GameRunner.Running() {
		return errors.New("game is already running")
	}
	err := gome.GameRunner.Init(game.CreateAdapters, settings)
	if err != nil {
		return err
	}
	err = gome.GameRunner.Loop(game.Initialize, game.Update, game.Draw)
	if err != nil {
		return err
	}
	return nil
}
