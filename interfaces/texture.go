package interfaces

import (
	"github.com/GomeBox/gome/primitives"
)

//Texture represents a drawable Texture
type Texture interface {
	//Unloader
	//Draw draws the part of the texture, defined by the source-rectangle to the part of the game window,
	//that is defined by the dest-rectangle. If source is nil, the whole texture is drawn
	Draw(source *primitives.Rectangle, dest primitives.RectangleF) error
	//Dimensions returns the textures dimensions
	Dimensions() primitives.Dimensions
}
