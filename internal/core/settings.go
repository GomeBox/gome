package core

import (
	"github.com/GomeBox/gome/adapters/graphics"
)

type Settings interface {
	WindowSettings() *graphics.WindowSettings
}
