package gome

import (
	"github.com/GomeBox/gome/adapters"
	g "github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/adapters/input"
	"github.com/GomeBox/gome/internal/core"
	"testing"
)

type gameRunnerMock struct {
	running       bool
	initializeCnt int
}

func (runner *gameRunnerMock) Initialize(initialize core.InitializeCallback, settings core.Settings) error {
	runner.initializeCnt++
	return nil
}

func (runner *gameRunnerMock) Loop(update core.UpdateCallBack, draw core.DrawCallback) error {
	return nil
}

func (runner *gameRunnerMock) Running() bool {
	return runner.running
}

func (runner *gameRunnerMock) AdapterSystem() adapters.System {
	return new(adapterSystemMock)
}

func (runner *gameRunnerMock) Quit() {

}

type adapterSystemMock struct {
}

func (adapterSystem *adapterSystemMock) Initialize() error {
	return nil
}

func (adapterSystem *adapterSystemMock) Update() error {
	return nil
}

func (adapterSystem *adapterSystemMock) Input() input.Port {
	return new(inputMock)
}

func (adapterSystem *adapterSystemMock) Graphics() g.Adapters {
	return new(graphicsMock)
}

type inputMock struct {
}

func (input *inputMock) Keyboard() input.Keyboard {
	return nil
}

func (input *inputMock) ControllerCount() int {
	return 0
}

func (input *inputMock) Controller(number int) (*input.Controller, error) {
	return nil, nil
}

type graphicsMock struct {
}

func (graphics *graphicsMock) TextureLoader() g.TextureLoader {
	return nil
}

func (graphics *graphicsMock) WindowAdapter() g.WindowAdapter {
	return nil
}

type gameMock struct {
}

func (game *gameMock) Initialize() (adapters.System, error) {
	return new(adapterSystemMock), nil
}

func (game *gameMock) Update(context Context) error {
	context.QuitGame()
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
	runner.running = true
	var err error
	err = gome.Run(game, settings)

	if err == nil {
		t.Errorf("game already running. Should be an error")
	}
}

func TestRun_InitializingGameRunner(t *testing.T) {
	gome := New()
	runner := new(gameRunnerMock)
	game := new(gameMock)
	settings := NewSettings()
	gome.ChangeGameRunner(runner)

	_ = gome.Run(game, settings)
	want := 1
	got := runner.initializeCnt
	if got != want {
		t.Errorf("was initialized %d times. Expected %d", got, want)
	}
}
