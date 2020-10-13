package adapters

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/adapters/input"
)

//System contains the systems that are used by the game to access the computers hardware like inputs and graphics
type System interface {
	//Initialize is called once at the start of the game and should initialize all systems
	Initialize() error
	//Update is called periodically by the game loop and should update all systems
	Update() error
	//Input returns the Inputadapter that is used to access keyboard, controllers, etc.
	Input() input.Port
	//Graphics returns the Graphicsadapter that is used to draw the game to the screen
	Graphics() graphics.Port
}
