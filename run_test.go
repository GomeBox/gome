package gome

import (
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/internal/core"
	"testing"
)

type gameRunnerMock struct {
	running bool
}

func (runner *gameRunnerMock) Initialize(initialize core.InitializeCallback, settings core.Settings) error {
	return nil
}

func (runner *gameRunnerMock) Loop(update core.UpdateCallBack, draw core.DrawCallback) chan error {
	return make(chan error, 1)
}

func (runner *gameRunnerMock) Running() bool {
	return runner.running
}

func (runner *gameRunnerMock) AdapterSystem() adapters.System {
	return nil
}

func (runner *gameRunnerMock) Quit() {

}

type gameMock struct {
}

func (game *gameMock) Initialize() (adapters.System, error) {
	return nil, nil
}

func (game *gameMock) Update(context Context) error {
	return nil
}

func (game *gameMock) Draw(context Context) error {
	return nil
}

func TestRun_ErrorIfAlreadyRunning(t *testing.T) {
	gome := New()
	runner := new(gameRunnerMock)
	game := new(gameMock)
	settings := NewSettings()
	gome.ChangeGameRunner(runner)

	err := gome.Run(game, settings)
	if err != nil {
		t.Errorf("game already running. Should be an error")
	}
}
