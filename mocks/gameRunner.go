package mocks

import (
	"github.com/GomeBox/gome/adapters"
	interfaces2 "github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/internal/game/interfaces"
)

type GameRunnerMock struct {
	IsRunning     bool
	InitializeCnt int
}

func (runner *GameRunnerMock) Loop(initialize interfaces.InitializeCallback, update interfaces.UpdateCallback, draw interfaces.DrawCallback) error {
	return nil
}

func (runner *GameRunnerMock) Init(createAdapters interfaces.CreateAdapters, settings interfaces2.Settings) error {
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
