package game

import "github.com/GomeBox/gome/adapters"

//Interface must be implemented to be run by Gome
type Interface interface {
	//Initialize is called by Gome once after calling Gome.Run
	Initialize() (adapters.System, error)
	//Update is called by Gome periodically to update the game's state
	Update(timeDelta float32, context Context) error
	//Draw is called by Gome periodically to draw the game to the screen
	Draw(timeDelta float32, context Context) error
}
