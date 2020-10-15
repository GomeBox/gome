package gome

import (
	"github.com/GomeBox/gome/adapters"
)

//Game must be implemented to be run by Gome
type Game interface {
	//Initialize is called by Gome once after calling Gome.Run
	Initialize() (adapters.System, error)
	//Update is called by Gome periodically to update the game's state
	Update(timeDelta float32, context Context) error
	//Draw is called by Gome periodically to draw the game to the screen
	Draw(timeDelta float32, context Context) error
}
