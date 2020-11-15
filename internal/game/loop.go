package game

type LoopData struct {
	gameSystem System
	update     func(timeDelta float32) (keepRunning bool, error error)
	draw       func(timeDelta float32) error
}

type Loop func(loopData *LoopData) error
