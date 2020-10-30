package game

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"time"
)

type loopData struct {
	screenPresenter graphics.ScreenPresenter
	update          func(timeDelta float32) (keepRunning bool, error error)
	draw            func(timeDelta float32) error
}

type loop func(loopData *loopData) error

func singleThreadedLoop(loopData *loopData) error {
	var loopStart time.Time
	var timeDelta float32
	var err error
	keepRunning := true
	for keepRunning {
		loopStart = time.Now()
		err = loopData.screenPresenter.Clear()
		if err != nil {
			return err
		}
		keepRunning, err = loopData.update(timeDelta)
		if err != nil {
			return err
		}
		err = loopData.draw(timeDelta)
		if err != nil {
			return err
		}
		err = loopData.screenPresenter.Present()
		if err != nil {
			return err
		}
		time.Sleep(time.Millisecond)
		timeDelta = float32(time.Since(loopStart).Milliseconds()) // / 1000000.0
		if timeDelta <= 0 {
			timeDelta = 1.0
		}
	}
	return nil
}
