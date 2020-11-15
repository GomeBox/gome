package game

import (
	graphicsAdapters "github.com/GomeBox/gome/internal/game/graphics"
	"testing"
)

func TestLoop_SingleThreadedLoop_ClearsGraphics(t *testing.T) {
	loopRuns := 3
	loopCnt := 0
	drawCnt := 0
	graphicsSystem := new(graphicsAdapters.SystemMock)
	var gameSystem System
	gameSystem = &systemMock{
		OnGraphics: func() graphicsAdapters.System {
			return graphicsSystem
		},
	}
	update := func(timeDelta float32) (keepRunning bool, error error) {
		loopCnt++
		if loopCnt >= loopRuns {
			return false, nil
		}
		return true, nil
	}
	draw := func(timeDelta float32) error {
		drawCnt++
		return nil
	}
	loopData := &LoopData{
		gameSystem: gameSystem,
		update:     update,
		draw:       draw,
	}
	err := SingleThreadedLoop(loopData)
	if err != nil {
		t.Fatal("singleThreadedLoop() returned unexpected error")
	}

	checkWasCalledAsExpected := func(funcName string, got, want int) {
		if got != want {
			t.Errorf("%s was not called expected number of times. Got %v, want %v", funcName, got, want)
		}
	}
	checkWasCalledAsExpected("update", loopCnt, loopRuns)
	checkWasCalledAsExpected("present", graphicsSystem.CallCntPresent, loopRuns)
	checkWasCalledAsExpected("clear", graphicsSystem.CallCntClear, loopRuns)
	checkWasCalledAsExpected("draw", drawCnt, loopRuns)
}
