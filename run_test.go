package gome

import (
	"github.com/GomeBox/gome/mocks"
	"testing"
)

func TestRun_ErrorIfAlreadyRunning(t *testing.T) {
	gome := New()
	runner := new(mocks.GameRunnerMock)
	game := new(mocks.GameMock)
	settings := NewSettings()
	gome.ChangeGameRunner(runner)
	runner.IsRunning = true
	var err error
	err = gome.Run(game, settings)

	if err == nil {
		t.Errorf("game already running. Should be an error")
	}
}

func TestRun_InitializingGameRunner(t *testing.T) {
	gome := New()
	runner := new(mocks.GameRunnerMock)
	game := new(mocks.GameMock)
	settings := NewSettings()
	gome.ChangeGameRunner(runner)

	_ = gome.Run(game, settings)
	want := 1
	got := runner.InitializeCnt
	if got != want {
		t.Errorf("was initialized %d times. Expected %d", got, want)
	}
}
