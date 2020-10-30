package game

import (
	"github.com/GomeBox/gome/audio"
	"github.com/GomeBox/gome/graphics"
	"github.com/GomeBox/gome/input"
)

//Context is passed to Interface.Update and Interface.Draw to access the game's systems, etc.
type Context interface {
	//Graphics returns the graphics adapter
	Graphics() graphics.System
	//Graphics returns the input adapter
	Input() input.System
	//Audio returns the audio adapter
	Audio() audio.System
}
