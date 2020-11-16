package graphics

import (
	"github.com/GomeBox/gome/interfaces"
)

type System interface {
	interfaces.Graphics
	OpenWindow(settings WindowSettings) error
	Clear() error
	Present() error
}
