package interfaces

import (
	"github.com/GomeBox/gome/primitives"
)

//Text represents a drawable Text
type Text interface {
	//Draw draws the text to the passed position of the game window
	Draw(position primitives.PointF) error
	//Dimensions returns the size of the text in pixels
	Dimensions() primitives.Dimensions
}
