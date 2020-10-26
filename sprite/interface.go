package sprite

import "github.com/GomeBox/gome/primitives"

//Interface describes a sprite
type Interface interface {
	//Dimensions returns the size of the Sprite
	Dimensions() primitives.Dimensions
	//Position returns the position of the sprite on the game window
	Position() primitives.PointF
	//SetPosition sets the position of the sprite on the game window
	SetPosition(x, y float32) error
	//Draw draws the sprite to it's Position
	Draw() error
}
