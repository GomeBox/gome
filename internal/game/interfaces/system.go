package interfaces

import (
	"github.com/GomeBox/gome/internal/audio"
	"github.com/GomeBox/gome/internal/graphics"
	"github.com/GomeBox/gome/internal/input"
)

type System interface {
	Initialize() error
	Update() error
	Graphics() *graphics.System
	Input() *input.System
	Audio() *audio.System
}
