package interfaces

import (
	gomeInterfaces "github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/internal/graphics"
)

type System interface {
	Initialize() error
	Update() error
	Graphics() graphics.System
	Context() gomeInterfaces.Context
	OpenGameWindow(settings graphics.WindowSettings) error
}
