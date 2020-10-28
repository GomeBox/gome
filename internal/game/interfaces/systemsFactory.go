package interfaces

import (
	audio "github.com/GomeBox/gome/internal/audio/interfaces"
	graphics "github.com/GomeBox/gome/internal/graphics/interfaces"
	input "github.com/GomeBox/gome/internal/input/interfaces"
)

type SystemsFactory interface {
	CreateGraphicsSystem() graphics.System
	CreateInputSystem() input.System
	CreateAudioSystem() audio.System
}
