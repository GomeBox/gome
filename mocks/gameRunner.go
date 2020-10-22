package mocks

import (
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/internal/game"
)

type GameRunnerMock struct {
	IsRunning     bool
	InitializeCnt int
}

func (runner *GameRunnerMock) Loop(initialize game.InitializeCallback, update game.UpdateCallback, draw game.DrawCallback) error {
	return nil
}

func (runner *GameRunnerMock) Init(createAdapters game.CreateAdapters, settings game.Settings) error {
	runner.InitializeCnt++
	return nil
}

func (runner *GameRunnerMock) Running() bool {
	return runner.IsRunning
}

func (runner *GameRunnerMock) AdapterSystem() adapters.System {
	return new(AdapterSystem)
}

func (runner *GameRunnerMock) Quit() {

}
