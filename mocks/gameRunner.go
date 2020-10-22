package mocks

import (
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/game"
	internalGame "github.com/GomeBox/gome/internal/game"
)

type GameRunnerMock struct {
	IsRunning     bool
	InitializeCnt int
}

func (runner *GameRunnerMock) Loop(initialize internalGame.InitializeCallback, update internalGame.UpdateCallback, draw internalGame.DrawCallback) error {
	return nil
}

func (runner *GameRunnerMock) Init(createAdapters internalGame.CreateAdapters, settings game.Settings) error {
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
