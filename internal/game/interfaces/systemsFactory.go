package interfaces

import (
	gomeInterfaces "github.com/GomeBox/gome/interfaces"
	graphics "github.com/GomeBox/gome/internal/graphics/interfaces"
	inputInterfaces "github.com/GomeBox/gome/internal/input/interfaces"
)

type SystemsFactory interface {
	CreateGraphicsSystem() graphics.System
	CreateInputSystem() inputInterfaces.System
	CreateAudioSystem() gomeInterfaces.Audio
}
