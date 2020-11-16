package internal

import (
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/internal/game"
)

//Start starts your game. You must provide the game instance itself and a function to create
//the hardware adapters. You can use a gome adapter package like sdl2 or provide your own adapters
func Start(yourGame interfaces.Game,
	adapterSystem adapters.System,
	systemsFactory SystemsFactory,
	gameLoop game.Loop) error {
	audioSystem := systemsFactory.CreateAudioSystem()
	graphicsSystem := systemsFactory.CreateGraphicsSystem()
	inputSystem := systemsFactory.CreateInputSystem()
	gameSystem := game.NewSystem(adapterSystem, graphicsSystem, audioSystem, inputSystem)
	runner := game.NewRunner(gameSystem, gameLoop)
	return runner.Run(yourGame)
}
