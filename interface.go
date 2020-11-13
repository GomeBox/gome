package gome

import (
	"github.com/GomeBox/gome/adapters"
)

//Interface must be implemented to be run by Gome
type Interface interface {
	//CreateAdapters is called by gome once to get the necessary adapters
	CreateAdapters() (adapters.System, error)
	//Initialize is called by Gome once after calling Gome.Run
	Initialize(context Context) error
	//Update is called by Gome periodically to update the game's state.
	//If this Function returns keepRunning = false the game quits
	Update(timeDelta float32, context Context) (keepRunning bool, error error)
	//Draw is called by Gome periodically to draw the game to the screen
	Draw(timeDelta float32, context Context) error
}
