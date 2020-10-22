package tests

import (
	game2 "github.com/GomeBox/gome/game"
	"github.com/GomeBox/gome/internal"
	"github.com/GomeBox/gome/mocks"
	"testing"
)

func TestRun_ErrorIfAlreadyRunning(t *testing.T) {
	gome := internal.GomeImpl{}
	runner := new(mocks.GameRunnerMock)
	game := new(mocks.GameMock)
	settings := game2.NewSettings()
	gome.GameRunner = runner
	runner.IsRunning = true
	var err error
	err = gome.Run(game, settings)

	if err == nil {
		t.Errorf("game already running. Should be an error")
	}
}

func TestRun_InitializingGameRunner(t *testing.T) {
	gome := internal.GomeImpl{}
	runner := new(mocks.GameRunnerMock)
	game := new(mocks.GameMock)
	settings := game2.NewSettings()
	gome.GameRunner = runner

	_ = gome.Run(game, settings)
	want := 1
	got := runner.InitializeCnt
	if got != want {
		t.Errorf("was initialized %d times. Expected %d", got, want)
	}
}
