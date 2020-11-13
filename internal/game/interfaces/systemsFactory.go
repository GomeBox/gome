package interfaces

import (
	gomeInterfaces "github.com/GomeBox/gome/interfaces"
	graphicsInterfaces "github.com/GomeBox/gome/internal/graphics/interfaces"
	inputInterfaces "github.com/GomeBox/gome/internal/input/interfaces"
)

type SystemsFactory interface {
	CreateGraphicsSystem() graphicsInterfaces.System
	CreateInputSystem() inputInterfaces.System
	CreateAudioSystem() gomeInterfaces.Audio
}
