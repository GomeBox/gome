package game

import (
	"github.com/GomeBox/gome/interfaces"
)

type Runner interface {
	Run(game interfaces.Game) error
}

type runner struct {
	context    interfaces.Context
	gameSystem System
	loop       Loop
}

func NewRunner(gameSystem System, gameLoop Loop) Runner {
	runner := runner{
		gameSystem: gameSystem,
		loop:       gameLoop,
	}
	return &runner
}

func (runner *runner) Run(game interfaces.Game) error {
	settings := newSettings()
	game.Setup(settings)
	err := runner.initialize(game, settings)
	if err != nil {
		return err
	}
	return runner.startLoop(game)
}

func (runner *runner) initialize(game interfaces.Game, settings *settings) error {
	err := runner.gameSystem.Initialize()
	if err != nil {
		return err
	}
	err = runner.gameSystem.OpenGameWindow(*settings.windowSettings)
	if err != nil {
		return err
	}
	err = game.Initialize(runner.gameSystem.Context())
	if err != nil {
		return err
	}
	return nil
}

func (runner *runner) startLoop(game interfaces.Game) error {
	loopData := runner.createLoopData(game)
	return runner.loop(loopData)
}

func (runner *runner) createLoopData(game interfaces.Game) *LoopData {
	return &LoopData{
		gameSystem: runner.gameSystem,
		update: func(timeDelta float32) (keepRunning bool, error error) {
			return runner.update(game, timeDelta)
		},
		draw: func(timeDelta float32) error {
			return game.Draw(timeDelta, runner.gameSystem.Context())
		},
	}
}

func (runner *runner) update(game interfaces.Game, timeDelta float32) (keepRunning bool, err error) {
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
