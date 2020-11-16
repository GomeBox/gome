package gome

import (
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/internal/game"
)

type CreateAdapters func() (adapters.System, error)

//Start starts your game. You must provide the game instance itself and a function to create
//the hardware adapters. You can use a gome adapter package like sdl2 or provide your own adapters
func Start(yourGame interfaces.Game,
	createAdapters CreateAdapters,
	systemsFactory SystemsFactory,
	gameLoop game.Loop) error {
	a, err := createAdapters()
	if err != nil {
		return err
	}
	audioSystem := systemsFactory.CreateAudioSystem(a.Audio())
	graphicsSystem := systemsFactory.CreateGraphicsSystem(a.Graphics())
	inputSystem := systemsFactory.CreateInputSystem(a.Input())
	gameSystem := game.NewSystem(a, graphicsSystem, audioSystem, inputSystem)
	runner := game.NewRunner(gameSystem, gameLoop)
	return runner.Run(yourGame)
}
