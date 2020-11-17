package graphics

import "github.com/GomeBox/gome/primitives"

//TextDrawer is the adapter that is used to draw text
type TextDrawer interface {
	//Draw renders the text in its original size to the defined position
	Draw(position primitives.PointF) error
	//DrawScaled draws the text into the defined destRect
	DrawScaled(destRect primitives.RectangleF) error
	//Dimensions returns the size of the text
	Dimensions() primitives.Dimensions
}
