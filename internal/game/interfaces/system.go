package interfaces

import (
	gomeInterfaces "github.com/GomeBox/gome/interfaces"
	graphics "github.com/GomeBox/gome/internal/graphics/interfaces"
)

type System interface {
	Initialize() error
	Update() error
	Graphics() graphics.System
	Context() gomeInterfaces.Context
	OpenGameWindow(settings gomeInterfaces.WindowSettings) error
}
