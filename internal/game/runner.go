package game

import (
	gomeInterfaces "github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/internal/game/interfaces"
)

type runner struct {
	context    gomeInterfaces.Context
	gameSystem interfaces.System
	loop       loop
}

func NewRunner(gameSystem interfaces.System, gameLoop loop) interfaces.Runner {
	runner := runner{
		gameSystem: gameSystem,
		loop:       gameLoop,
	}
	return &runner
}

func (runner *runner) Run(game gomeInterfaces.Game) error {
	settings := NewSettings()
	game.Setup(settings)
	err := runner.initialize(game, settings)
	if err != nil {
		return err
	}
	return runner.startLoop(game)
}

func (runner *runner) initialize(game gomeInterfaces.Game, settings gomeInterfaces.Settings) error {
	err := runner.gameSystem.Initialize()
	if err != nil {
		return err
	}
	err = runner.gameSystem.OpenGameWindow(settings.WindowSettings())
	if err != nil {
		return err
	}
	err = game.Initialize(runner.gameSystem.Context())
	if err != nil {
		return err
	}
	return nil
}

func (runner *runner) startLoop(game gomeInterfaces.Game) error {
	loopData := runner.createLoopData(game)
	return runner.loop(loopData)
}

func (runner *runner) createLoopData(game gomeInterfaces.Game) *loopData {
	return &loopData{
		gameSystem: runner.gameSystem,
		update: func(timeDelta float32) (keepRunning bool, error error) {
			return runner.update(game, timeDelta)
		},
		draw: func(timeDelta float32) error {
			return game.Draw(timeDelta, runner.gameSystem.Context())
		},
	}
}

func (runner *runner) update(game gomeInterfaces.Game, timeDelta float32) (keepRunning bool, err error) {
	err = runner.gameSystem.Update()
	if err != nil {
		return false, err
	}
	keepRunning, err = game.Update(timeDelta, runner.gameSystem.Context())
	if err != nil {
		return false, err
	}
	return keepRunning, nil
}
