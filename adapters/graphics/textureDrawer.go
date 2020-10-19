package graphics

import "github.com/GomeBox/gome/primitives"

//TextureDrawer is the adapter that draws a texture to the screen
type TextureDrawer interface {
	//Draw draws the part of the texture, that is defined by source to the location defined by dest
	//Pass source == nil to draw the whole texture
	Draw(source, dest *primitives.Rectangle) error
	//DrawF draws the part of the texture, that is defined by source to the location defined by dest
	//Pass source == nil to draw the whole texture
	DrawF(source *primitives.Rectangle, dest *primitives.RectangleF) error
}
