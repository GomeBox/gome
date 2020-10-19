package mocks

import (
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/internal/core"
)

type GameRunnerMock struct {
	IsRunning     bool
	InitializeCnt int
}

func (runner *GameRunnerMock) Initialize(initialize core.InitializeCallback, settings core.Settings) error {
	runner.InitializeCnt++
	return nil
}

func (runner *GameRunnerMock) Loop(update core.UpdateCallBack, draw core.DrawCallback) error {
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
