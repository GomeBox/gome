package internal

import (
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/internal/game"
)

func Run(gameImpl interfaces.Game, adapterSystem adapters.System) error {
	loop := game.SingleThreadedLoop
	systemsFactory := game.NewSystemsFactory(adapterSystem)
	gameSystem := game.NewSystem(adapterSystem, systemsFactory)
	runner := game.NewRunner(gameSystem, loop)
	return runner.Run(gameImpl)
}
