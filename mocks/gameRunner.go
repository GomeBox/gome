package mocks

import (
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/internal/core"
)

type GameRunnerMock struct {
	IsRunning     bool
	InitializeCnt int
}

func (runner *GameRunnerMock) Loop(initialize core.InitializeCallback, update core.UpdateCallback, draw core.DrawCallback) error {
	return nil
}

func (runner *GameRunnerMock) Init(createAdapters core.CreateAdapters, settings core.Settings) error {
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
