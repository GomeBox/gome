package interfaces

import (
	gomeInterfaces "github.com/GomeBox/gome/interfaces"
	graphics "github.com/GomeBox/gome/internal/graphics/interfaces"
	input "github.com/GomeBox/gome/internal/input/interfaces"
)

type System interface {
	Initialize() error
	Update() error
	Input() input.System
	Audio() gomeInterfaces.Audio
	Graphics() graphics.System
}
