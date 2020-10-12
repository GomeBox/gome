package adapters

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/adapters/input"
)

type System interface {
	Initialize() error
	Update() error
	Input() input.Interface
	Graphics() graphics.Interface
}
