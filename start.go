package gome

import (
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/internal/audio"
	"github.com/GomeBox/gome/internal/game"
	"github.com/GomeBox/gome/internal/graphics"
	"github.com/GomeBox/gome/internal/input"
)

type CreateAdapters func() (adapters.System, error)

//Start starts your game. You must provide the game instance itself and a function to create
//the hardware adapters. You can use a gome adapter package like sdl2 or provide your own adapters
func Start(yourGame interfaces.Game, createAdapters CreateAdapters) error {
	a, err := createAdapters()
	if err != nil {
		return err
	}
	audioSystem := audio.NewSystem(a.Audio())
	graphicsSystem := graphics.NewSystem(a.Graphics())
	inputSystem := input.NewSystem(a.Input())
	gameSystem := game.NewSystem(a, graphicsSystem, audioSystem, inputSystem)
	gameLoop := game.SingleThreadedLoop
	runner := game.NewRunner(gameSystem, gameLoop)
	return runner.Run(yourGame)
}
