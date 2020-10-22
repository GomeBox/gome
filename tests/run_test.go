package tests

import (
	"github.com/GomeBox/gome/internal"
	"github.com/GomeBox/gome/mocks"
	"testing"
)

func TestRun_ErrorIfAlreadyRunning(t *testing.T) {
	gome := internal.GomeImpl{}
	runner := new(mocks.GameRunnerMock)
	game := new(mocks.GameMock)
	gome.GameRunner = runner
	runner.IsRunning = true
	var err error
	err = gome.Run(game)

	if err == nil {
		t.Errorf("game already running. Should be an error")
	}
}

func TestRun_InitializingGameRunner(t *testing.T) {
	gome := internal.GomeImpl{}
	runner := new(mocks.GameRunnerMock)
	game := new(mocks.GameMock)
	gome.GameRunner = runner

	_ = gome.Run(game)
	want := 1
	got := runner.InitializeCnt
	if got != want {
		t.Errorf("was initialized %d times. Expected %d", got, want)
	}
}
