package game

import (
	"github.com/GomeBox/gome/internal/game/interfaces"
	"github.com/GomeBox/gome/internal/game/mocks"
	"github.com/GomeBox/gome/internal/graphics"
	graphicsInterfaces "github.com/GomeBox/gome/internal/graphics/interfaces"
	"testing"
)

func TestLoop_SingleThreadedLoop_ClearsGraphics(t *testing.T) {
	loopRuns := 3
	loopCnt := 0
	drawCnt := 0
	graphicsSystem := new(graphics.SystemMock)
	var gameSystem interfaces.System
	gameSystem = &mocks.System{
		OnGraphics: func() graphicsInterfaces.System {
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
	loopData := &loopData{
		gameSystem: gameSystem,
		update:     update,
		draw:       draw,
	}
	err := singleThreadedLoop(loopData)
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
