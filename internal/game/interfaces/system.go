package interfaces

import (
	audio "github.com/GomeBox/gome/internal/audio/interfaces"
	graphics "github.com/GomeBox/gome/internal/graphics/interfaces"
	input "github.com/GomeBox/gome/internal/input/interfaces"
)

type System interface {
	Initialize() error
	Update() error
	Graphics() graphics.System
	Input() input.System
	Audio() audio.System
}
