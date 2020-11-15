package game

import "time"

func SingleThreadedLoop(loopData *LoopData) error {
	var loopStart time.Time
	var timeDelta float32
	var err error
	keepRunning := true
	for keepRunning {
		loopStart = time.Now()
		err = loopData.gameSystem.Graphics().Clear()
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
		err = loopData.gameSystem.Graphics().Present()
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
