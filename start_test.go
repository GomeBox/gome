package gome

import (
	"testing"
)

func TestGome_NewCreatesInstance(t *testing.T) {
	//gome := New()
	//if gome == nil {
	//	t.Errorf("New() returned nil")
	//}
}

//func TestGome_ErrorIfAlreadyRunning(t *testing.T) {
//	gome := internal.gomeImpl{}
//	runner := new(mocks.GameRunnerMock)
//	game := new(mocks.GameMock)
//	gome.gameRunner = runner
//	runner.IsRunning = true
//	var err error
//	err = gome.Run(game)
//
//	if err == nil {
//		t.Errorf("game already running. Should be an error")
//	}
//}
//
//func TestGome_InitializingGameRunner(t *testing.T) {
//	gome := internal.gomeImpl{}
//	runner := new(mocks.GameRunnerMock)
//	game := new(mocks.GameMock)
//	gome.gameRunner = runner
//
//	_ = gome.Run(game)
//	want := 1
//	got := runner.InitializeCnt
//	if got != want {
//		t.Errorf("was initialized %d times. Expected %d", got, want)
//	}
//}
