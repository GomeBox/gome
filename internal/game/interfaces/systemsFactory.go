package interfaces

import (
	audioInterfaces "github.com/GomeBox/gome/internal/audio/interfaces"
	graphicsInterfaces "github.com/GomeBox/gome/internal/graphics/interfaces"
	inputInterfaces "github.com/GomeBox/gome/internal/input/interfaces"
)

type SystemsFactory interface {
	CreateGraphicsSystem() graphicsInterfaces.System
	CreateInputSystem() inputInterfaces.System
	CreateAudioSystem() audioInterfaces.System
}
